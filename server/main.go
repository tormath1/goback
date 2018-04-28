package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	"github.com/otiai10/copy"
	pb "github.com/tormath1/goback/server/proto"
)

var docker *client.Client

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

	log.Fatal(grpcServer.Serve(listener))
}

type server struct{ docker *client.Client }

func (s *server) SaveVolume(ctx context.Context, in *pb.SaveVolumeRequest) (*pb.Error, error) {

	src := in.VolumeName
	dst := in.Destination

	volumes, err := s.docker.VolumeList(ctx, filters.Args{})
	if err != nil {
		msg := fmt.Sprintf("unable to list volumes: %v", err)
		log.Print(msg)
		return &pb.Error{Message: msg}, err
	}

	for _, volume := range volumes.Volumes {
		if volume.Name == src {
			if err = copy.Copy(volume.Mountpoint, fmt.Sprintf("%s/_data", dst)); err != nil {
				msg := fmt.Sprintf("unable to copy volume: %s on %s: %v", volume.Name, dst, err)
				log.Print(msg)
				return &pb.Error{Message: msg}, err
			}
		}
	}
	return &pb.Error{}, nil
}
