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

/// receive usecase and add it to local data base
func (s *server) Push(ctx context.Context, usecase *protocol.Usecase) (*protocol.Usecase, error) {
	klog.Info(fmt.Sprintf("Received: %v", usecase))
	return usecase, nil
}

//// receive config data from controller
func StartGrpcServer(done chan (struct{})) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", config_port))
	if err != nil {
		klog.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	protocol.RegisterConWireServer(s, &server{})
	klog.Info(fmt.Sprintf("config server listening at %v", lis.Addr()))
	if err := s.Serve(lis); err != nil {
		close(done)
		log.Fatalf("failed to serve: %v", err)
	}
}
