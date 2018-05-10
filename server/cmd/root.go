package cmd

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/boltdb/bolt"
	"github.com/grpc-ecosystem/go-grpc-prometheus"

	"google.golang.org/grpc/credentials"

	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	"github.com/otiai10/copy"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/robfig/cron"
	"github.com/spf13/cobra"
	pb "github.com/tormath1/goback/server/proto"
	"google.golang.org/grpc"
)

var (
	docker                    *client.Client
	chronoTable               *cron.Cron
	certFile, keyFile, rootCA string
	rootCmd                   = &cobra.Command{
		Use:   "manager ",
		Short: "main server to handle goback requests",
		Run: func(cmd *cobra.Command, args []string) {
			serve(cmd, args)
		},
	}
	nbCron = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "number_cron_job",
		Help: "Number of cron job",
	})
	db *bolt.DB
)

func init() {
	rootCmd.Flags().StringVarP(&certFile, "cert-file", "", "", "absolute path to a certificat file")
	rootCmd.Flags().StringVarP(&keyFile, "key-file", "", "", "absolute path to a certificat key file")
	rootCmd.Flags().StringVarP(&rootCA, "root-ca", "", "", "absolute path to root authority (*.pem / *.crt) file")
	prometheus.Register(nbCron)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("unable to run root cmd: %v", err)
	}
}

func serve(cmd *cobra.Command, args []string) {
	docker, err := client.NewEnvClient()
	if err != nil {
		log.Fatalf("unable to connect to docker daemon: %v", err)
	}
	listener, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("unable to listen on :12800: %v", err)
	}

	http.Handle("/metrics", promhttp.Handler())
	go http.ListenAndServe(":8080", nil)

	var grpcServer *grpc.Server
	if certFile != "" && keyFile != "" {
		grpcServer = grpc.NewServer(
			getGrpcCreds(),
			grpc.StreamInterceptor(grpc_prometheus.StreamServerInterceptor),
			grpc.UnaryInterceptor(grpc_prometheus.UnaryServerInterceptor),
		)
	} else {
		grpcServer = grpc.NewServer(
			grpc.StreamInterceptor(grpc_prometheus.StreamServerInterceptor),
			grpc.UnaryInterceptor(grpc_prometheus.UnaryServerInterceptor),
		)
	}

	grpc_prometheus.Register(grpcServer)
	pb.RegisterManagerServer(grpcServer, &server{docker})

	chronoTable = cron.New()
	chronoTable.Start()

	db, err = bolt.Open("/tmp/backups.db", 0600, nil)
	defer db.Close()
	if err != nil {
		log.Fatalf("unable to open database: %v", err)
	}
	if err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("entries"))
		return err
	}); err != nil {
		log.Fatalf("unable to create bucket `entries`: %v", err)
	}

	log.Fatal(grpcServer.Serve(listener))
}

func getGrpcCreds() grpc.ServerOption {
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		log.Fatalf("unable to load certificate: %v", err)
	}
	if rootCA != "" {
		certPool := x509.NewCertPool()
		ca, err := ioutil.ReadFile(rootCA)
		if err != nil {
			log.Fatalf("unable to load root authority: %v", err)
		}
		if ok := certPool.AppendCertsFromPEM(ca); !ok {
			log.Fatalf("unable to append root authority to cert pool: %v", err)
		}
		return grpc.Creds(credentials.NewTLS(&tls.Config{
			ClientAuth:   tls.RequireAndVerifyClientCert,
			Certificates: []tls.Certificate{cert},
			ClientCAs:    certPool,
		}))
	}
	return grpc.Creds(credentials.NewTLS(&tls.Config{
		ClientAuth:   tls.RequireAndVerifyClientCert,
		Certificates: []tls.Certificate{cert},
	}))
}

type server struct{ docker *client.Client }

func (s *server) Restore(ctx context.Context, in *pb.RestoreVolumeRequest) (*pb.Empty, error) {
	return new(pb.Empty), fmt.Errorf("not implemented")
}

func (s *server) ListEntries(ctx context.Context, in *pb.Empty) (*pb.EntriesList, error) {
	out := &pb.EntriesList{}
	if err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("entries"))
		return b.ForEach(func(key, value []byte) error {
			out.Entries = append(out.Entries, &pb.Entry{Volume: string(key), Cron: string(value)})
			return nil
		})
	}); err != nil {
		log.Printf("unable to get entries: %v", err)
		return new(pb.EntriesList), err
	}
	return out, nil
}

func (s *server) SaveVolume(ctx context.Context, in *pb.SaveVolumeRequest) (*pb.Empty, error) {

	src := in.VolumeName
	dst := in.Destination

	mountpoint, err := getMountpoint(src, s.docker)
	if err != nil {
		return &pb.Empty{}, err
	}

	err = save(src, dst, mountpoint)
	if err != nil {
		log.Printf("unable to save volume: %v", err)
	}
	return &pb.Empty{}, err
}

func (s *server) ScheduleSaving(ctx context.Context, in *pb.ScheduleSavingRequest) (*pb.Empty, error) {

	mountpoint, err := getMountpoint(in.Volume.VolumeName, s.docker)
	if err != nil {
		return &pb.Empty{}, err
	}
	job := func() { save(in.Volume.VolumeName, in.Volume.Destination, mountpoint) }
	err = chronoTable.AddFunc(in.Schedule, job)
	if err != nil {
		log.Printf("unable to add entry to chrono table: %v", err)
	}

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("entries"))
		cronEntry := b.Get([]byte(in.Volume.VolumeName))
		if cronEntry == nil {
			nbCron.Inc()
		}
		return nil
	})
	if err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("entries"))
		return b.Put([]byte(in.Volume.VolumeName), []byte(in.Schedule))
	}); err != nil {
		log.Printf("unable to save cron in db: %v", err)
	}
	return &pb.Empty{}, err
}

func save(src, dst, mountpoint string) error {

	timestamp := time.Now().Format(time.RFC3339)

	if err := copy.Copy(mountpoint, fmt.Sprintf("%s/%s-%s/_data", dst, src, timestamp)); err != nil {
		return err
	}
	return nil
}

func getMountpoint(src string, cli *client.Client) (string, error) {
	volumes, err := cli.VolumeList(context.Background(), filters.Args{})
	if err != nil {
		return "", err
	}
	for _, volume := range volumes.Volumes {
		if volume.Name == src {
			return volume.Mountpoint, nil
		}
	}
	return "", fmt.Errorf("volume %s does not exist", src)
}
