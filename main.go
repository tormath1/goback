package main

import (
	"context"
	"log"
	"os"

	pb "github.com/tormath1/goback/server/proto"
	"google.golang.org/grpc"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		log.Fatal("usage: ./goback <command> <arguments>")
	}

	client, err := grpc.Dial("localhost:12800", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("unable to connect to localhost:12800: %v", err)
	}
	defer client.Close()

	manager := pb.NewManagerClient(client)

	switch args[0] {
	case "save":
		saveVolume(manager, args[1:]...)
	default:
		log.Println("list of commands: \nsave <src> <dst>")
	}
}

func saveVolume(manager pb.ManagerClient, args ...string) error {
	log.Printf("save: %s on %s", args[0], args[1])
	_, err := manager.SaveVolume(context.Background(), &pb.SaveVolumeRequest{
		VolumeName:  args[0],
		Destination: args[1],
	})
	if err != nil {
		log.Fatalf("unable to save volume: %v", err)
	}
	return err
}
