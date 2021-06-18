package main

import (
	"fmt"
	"github.com/Hristiyan-Bonev/golang-course/sort/gen"
	"google.golang.org/grpc"
	"log"
	"net"
	"google.golang.org/grpc/reflection"
)

const serverAddress = "localhost:10001"

func main() {

	server, lis := NewFulfillmentServer()

	fmt.Printf("gRPC server started on port")

	if err := server.Serve(lis); err != nil {
		log.Fatalf("cannot serve on port %s. Reason: %s", serverAddress, err)
	}

}

func NewFulfillmentServer() (*grpc.Server, net.Listener) {

	lis, err := net.Listen("tcp", serverAddress)

	if err != nil {
		log.Fatalf("cannot listen : %s", err)
	}

	grpcServer := grpc.NewServer()
	gen.RegisterFulfillmentServer(grpcServer, NewFulfillmentService())
	reflection.Register(grpcServer)
	return grpcServer, lis
}
