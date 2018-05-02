package cmd

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"

	"google.golang.org/grpc/credentials"

	"github.com/spf13/viper"

	pb "github.com/tormath1/goback/server/proto"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var rootCmd = &cobra.Command{
	Use:   "goback ",
	Short: "schedule your Docker volumes backups with goback",
}

var manager pb.ManagerClient

func init() {
	viper.SetConfigName("goback")
	viper.AddConfigPath(".")
	viper.AddConfigPath("/etc/goback")
	viper.AddConfigPath("local/")

	viper.SetDefault("address", "localhost:9090")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("unable to read config: %v", err)
	}

	client, err := grpc.Dial(viper.GetString("address"), getGrpcOptions(
		viper.GetString("tls.cert_file"),
		viper.GetString("tls.key_file"),
		viper.GetString("tls.root_cert"),
		viper.GetString("tls.server_name"),
	))
	//client, err := grpc.Dial(viper.GetString("address"), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("unable to connect to %s %v", viper.GetString("address"), err)
	}
	manager = pb.NewManagerClient(client)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("unable to execute command: %v", err)
	}
}

func getGrpcOptions(crt, key, root, serverName string) grpc.DialOption {
	if crt != "" && key != "" && serverName != "" {
		cert, err := tls.LoadX509KeyPair(crt, key)
		if err != nil {
			log.Fatalf("unable to load cert files: %v", err)
		}
		if root != "" {
			certPool := x509.NewCertPool()
			ca, err := ioutil.ReadFile(root)
			if err != nil {
				log.Fatalf("unable to read cert authority: %v", err)
			}
			if ok := certPool.AppendCertsFromPEM(ca); !ok {
				log.Fatalf("unable to append root authority to cert pool: %v", err)
			}
			return grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{
				Certificates: []tls.Certificate{cert},
				RootCAs:      certPool,
				ServerName:   serverName,
			}))
		}
		return grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{
			Certificates: []tls.Certificate{cert},
			ServerName:   serverName,
		}))
	}
	return grpc.WithInsecure()
}
