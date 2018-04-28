package main

import (
	"context"
	"fmt"
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
	case "schedule":
		schedule(manager, args[1:]...)
	default:
		log.Println("list of commands: \nsave <src> <dst>")
	}
}

func schedule(manager pb.ManagerClient, args ...string) {
	switch args[0] {
	case "list":
		entriesList(manager)
	default:
		scheduleVolume(manager, args...)
	}
}

func saveVolume(manager pb.ManagerClient, args ...string) {
	_, err := manager.SaveVolume(context.Background(), &pb.SaveVolumeRequest{
		VolumeName:  args[0],
		Destination: args[1],
	})
	if err != nil {
		log.Fatalf("unable to save volume: %v", err)
	}
}

func scheduleVolume(manager pb.ManagerClient, args ...string) {
	_, err := manager.ScheduleSaving(context.Background(), &pb.ScheduleSavingRequest{
		Schedule: args[2],
		Volume: &pb.SaveVolumeRequest{
			VolumeName:  args[0],
			Destination: args[1],
		},
	})
	if err != nil {
		log.Fatalf("unable to get cron entries: %v", err)
	}
}

func entriesList(manager pb.ManagerClient) {
	entries, err := manager.ListEntries(context.Background(), &pb.Empty{})
	if err != nil {
		log.Fatalf("unable to get cron entries: %v", err)
	}
	for _, entry := range entries.Entries {
		fmt.Println(entry)
	}
}
