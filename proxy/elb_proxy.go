package main

import (
	"github.com/gnunu/elb/usecase"
)

const (
	config_port  = 55554
	request_port = 55555
)

var (
	usecases *usecase.UsecaseSet
)

func main() {
	done := make(chan struct{})
	usecases = usecase.NewUsecaseSet()
	go StartGrpcServer(done)
	go StartHttpServer(done)
	<-done
}
