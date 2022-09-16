package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/gnunu/elb/protocol"

	"google.golang.org/grpc"
	"k8s.io/klog"
)

type server struct {
	protocol.UnimplementedConWireServer
}

var (
	port = 55555
)

func (s *server) Push(ctx context.Context, usecase *protocol.Usecase) (*protocol.Usecase, error) {
	klog.Info(fmt.Sprintf("Received: %v", usecase))
	return usecase, nil
}

func StartGrpcServer() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		klog.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	protocol.RegisterConWireServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
