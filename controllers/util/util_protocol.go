package util

import (
	"context"
	"fmt"
	"time"

	"github.com/gnunu/elb/protocol"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"k8s.io/klog"
)

var (
	addr = "localhost:55554"
)

func SendUsecase(usecase *protocol.Usecase) {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		klog.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := protocol.NewConWireClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err = c.Push(ctx, usecase)
	if err != nil {
		klog.Fatalf("could not push: %v", err)
	}
	klog.Info(fmt.Sprintf("Send: %v", usecase))
}
