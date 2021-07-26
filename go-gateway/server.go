package main

import (
	"context"
	demo "go-gateway-protos/demo"
	"log"
	"net"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct {
	demo.UnimplementedDemoServiceServer
}

func (s *server) Echo(ctx context.Context, msg *demo.StringMessage) (*demo.StringMessage, error) {
	log.Printf("Received Echo:", msg)
	return msg, nil
}

func (s *server) SayHello(ctx context.Context, in *demo.HelloRequest) (*demo.HelloReply, error) {
	log.Printf("Received SayHello: %v", in.GetName())
	return &demo.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	demo.RegisterDemoServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
