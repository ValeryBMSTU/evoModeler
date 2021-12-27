package main

import (
	"github.com/ValeryBMSTU/evoModeler/internal/auth"
	"github.com/ValeryBMSTU/evoModeler/internal/bl"
	"github.com/ValeryBMSTU/evoModeler/internal/da"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":5301")

	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)

	serverDa, err := da.CreateDa()
	if err != nil {
		log.Fatal(err)
	}
	serverBl := bl.Bl{serverDa}
	if err != nil {
		log.Fatal(err)
	}

	auth.RegisterAuthServer(grpcServer, &auth.Server{serverBl})
	grpcServer.Serve(listener)
}
