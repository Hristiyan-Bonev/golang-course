package main

import (
	"github.com/Hristiyan-Bonev/golang-course/sort/gen"
	"google.golang.org/grpc"
	"log"
	"net"
)

const serverAddress := "localhost"
const serverPort := ":10001"

func main() {

	server, lis := NewFulfillmentServer()
}

func NewFulfillmentServer() (*grpc.Server, net.Listener){
	address := net.JoinHostPort(serverAddress, serverPort)
	lis, err := net.Listen("tcp", address)

	if err != nil {
		log.Fatalf("cannot listen ")
	}

	grpcServer := grpc.NewServer()
	gen.RegisterFulfillmentServer(grpcServer, NewFulfillmentService())
	return grpcServer, lis
}