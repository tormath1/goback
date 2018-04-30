package cmd

import (
	"log"

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
	client, err := grpc.Dial("localhost:12800", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("unable to connect to localhost:12800: %v", err)
	}
	manager = pb.NewManagerClient(client)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("unable to execute command: %v", err)
	}
}
