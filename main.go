package main

import (
	"log"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		log.Fatal("usage: ./goback <command> <arguments>")
	}

	switch args[0] {
	case "save":
		save(args[1:]...)
	default:
		log.Println("list of commands: \nsave <src> <dst>")
	}
}

func save(arguments ...string) {
	if len(arguments) < 2 {
		log.Fatalf("usage: ./goback save <src> <dst>")
	}

	src := arguments[0]
	dst := arguments[1]

	log.Printf("save %s to %s", src, dst)
}
