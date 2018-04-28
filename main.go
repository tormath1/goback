package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	"github.com/otiai10/copy"
)

func main() {
	docker, err := client.NewEnvClient()
	if err != nil {
		log.Fatalf("unable to connect to docker daemon: %v", err)
	}

	ctx := context.Background()
	args := os.Args[1:]
	if len(args) < 1 {
		log.Fatal("usage: ./goback <command> <arguments>")
	}

	switch args[0] {
	case "save":
		save(ctx, docker, args[1:]...)
	default:
		log.Println("list of commands: \nsave <src> <dst>")
	}
}

func save(ctx context.Context, cli *client.Client, arguments ...string) {
	if len(arguments) < 2 {
		log.Fatalf("usage: ./goback save <src> <dst>")
	}

	src := arguments[0]
	dst := arguments[1]

	log.Printf("save %s to %s", src, dst)
	volumes, err := cli.VolumeList(ctx, filters.Args{})
	if err != nil {
		log.Fatalf("unable to list volumes: %v", err)
	}

	for _, volume := range volumes.Volumes {
		if volume.Name == src {
			if err = copy.Copy(volume.Mountpoint, fmt.Sprintf("%s/_data", dst)); err != nil {
				log.Fatalf("unable to copy volume: %s on %s: %v", volume.Name, dst, err)
			}
		}
	}
}
