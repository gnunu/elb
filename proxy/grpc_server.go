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
// endpoint format: {node-name:pod-name:ip:port}*
func (s *server) Push(ctx context.Context, u *protocol.Usecase) (*protocol.Response, error) {
	klog.Infof("Received: %v", u)
	devstr := u.Devices
	epstr := u.Endpoints
	devlist := strings.Split(devstr, ",")
	epstrlist := strings.Split(epstr, ",")
	eplist := []usecase.Endpoint{}
	for _, ep := range epstrlist {
		es := strings.Split(ep, ":")
		if len(es) != 4 {
			klog.Infof("not enough info for an endpoint: %v", es)
		}
		e := usecase.Endpoint{Node: es[0], Pod: es[1], Addr: es[2], Port: es[4]}
		eplist = append(eplist, e)
	}

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
	klog.Infof("config server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		close(done)
		log.Fatalf("failed to serve: %v", err)
	}
}
