package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/gnunu/elb/protocol"
	"github.com/gnunu/elb/usecase"

	"google.golang.org/grpc"
	"k8s.io/klog"
)

type server struct {
	protocol.UnimplementedConWireServer
}

/// receive usecase and add it to local data base
func (s *server) Push(ctx context.Context, u *protocol.Usecase) (*protocol.Response, error) {
	klog.Info(fmt.Sprintf("Received: %v", u))
	devstr := u.Devices
	epstr := u.Endpoints
	devlist := strings.Split(devstr, ",")
	eplist := strings.Split(epstr, ",")

	u2 := usecase.NewUsecase(u.Name, devlist, u.Policy, eplist)
	usecases.Update(u2)
	usecases.List()
	return &protocol.Response{Result: "success"}, nil
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
