package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"

	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	"github.com/otiai10/copy"
	"github.com/robfig/cron"
	pb "github.com/tormath1/goback/server/proto"
)

var docker *client.Client
var chronoTable *cron.Cron

func main() {
	docker, err := client.NewEnvClient()
	if err != nil {
		log.Fatalf("unable to connect to docker daemon: %v", err)
	}
	listener, err := net.Listen("tcp", ":12800")
	if err != nil {
		log.Fatalf("unable to listen on :12800: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterManagerServer(grpcServer, &server{docker})

	chronoTable = cron.New()
	log.Fatal(grpcServer.Serve(listener))
}

type server struct{ docker *client.Client }

func (s *server) ListEntries(ctx context.Context, in *pb.Empty) (*pb.EntriesList, error) {
	out := &pb.EntriesList{}
	entries := chronoTable.Entries()
	for _, entry := range entries {
		out.Entries = append(out.Entries, &pb.Entry{Next: entry.Next.String()})
	}

	return out, nil
}

func (s *server) SaveVolume(ctx context.Context, in *pb.SaveVolumeRequest) (*pb.Error, error) {
	out := &pb.Error{Code: 200}

	src := in.VolumeName
	dst := in.Destination

	err := save(src, dst, s.docker)
	if err != nil {
		log.Printf("unable to save volume: %v", err)
		out.Code = 500
		out.Message = err.Error()
	}
	return out, err
}

func (s *server) ScheduleSaving(ctx context.Context, in *pb.ScheduleSavingRequest) (*pb.Error, error) {

	out := &pb.Error{Code: 200}

	job := func() { save(in.Volume.VolumeName, in.Volume.Destination, docker) }
	err := chronoTable.AddFunc(in.Schedule, job)
	if err != nil {
		log.Printf("unable to add entry to chrono table: %v", err)
		out.Code = 500
		out.Message = err.Error()
	}
	return out, err
}

func save(src, dst string, cli *client.Client) error {

	volumes, err := cli.VolumeList(context.Background(), filters.Args{})
	if err != nil {
		return err
	}

	timestamp := time.Now().Format(time.RFC3339)

	for _, volume := range volumes.Volumes {
		if volume.Name == src {
			if err = copy.Copy(volume.Mountpoint, fmt.Sprintf("%s/%s-%s/_data", dst, volume.Name, timestamp)); err != nil {
				return err
			}
		}
	}
	return err
}
