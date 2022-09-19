package main

const (
	config_port  = 55554
	request_port = 55555
)

func main() {
	done := make(chan struct{})
	go StartGrpcServer()
	go StartHttpServer(done)
	<-done
}
