package main

import (
	"log"
	"net"

	"github.com/varasu/grpc-k8s-example/gservice/internal/server"
)

const (
	port = ":8080"
)

func main() {
	log.Printf("Starting listening on port %s\n", port)

	l, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("%v", err)
	}

	log.Printf("Listening on %s", port)
	srv := server.NewGRPCServer()
	if err := srv.Serve(l); err != nil {
		log.Fatalf("%v", err)
	}
}
