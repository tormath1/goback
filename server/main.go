package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	"github.com/otiai10/copy"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/robfig/cron"
	pb "github.com/tormath1/goback/server/proto"
	"google.golang.org/grpc"
)

var docker *client.Client
var chronoTable *cron.Cron
var entries map[string]string

var nbCron = prometheus.NewCounter(prometheus.CounterOpts{
	Name: "number_cron_job",
	Help: "Number of cron job",
})

func init() {
	prometheus.Register(nbCron)
}

func main() {
	docker, err := client.NewEnvClient()
	if err != nil {
		log.Fatalf("unable to connect to docker daemon: %v", err)
	}
	listener, err := net.Listen("tcp", ":12800")
	if err != nil {
		log.Fatalf("unable to listen on :12800: %v", err)
	}

	http.Handle("/metrics", promhttp.Handler())
	go http.ListenAndServe(":8080", nil)

	grpcServer := grpc.NewServer()
	pb.RegisterManagerServer(grpcServer, &server{docker})

	chronoTable = cron.New()
	chronoTable.Start()
	entries = make(map[string]string)
	log.Fatal(grpcServer.Serve(listener))
}

type server struct{ docker *client.Client }

func (s *server) ListEntries(ctx context.Context, in *pb.Empty) (*pb.EntriesList, error) {
	out := &pb.EntriesList{}
	for key, value := range entries {
		out.Entries = append(out.Entries, &pb.Entry{Volume: key, Cron: value})
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
	if _, ok := entries[in.Volume.VolumeName]; !ok {
		nbCron.Inc()
	}
	entries[in.Volume.VolumeName] = in.Schedule
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
