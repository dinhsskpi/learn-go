package main

import (
	"context"
	"log"
	"net"
	"time"

	helloworld "protos/helloworld"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct {
	helloworld.UnimplementedGreeterServiceServer
}

func (s *server) SayHello(ctx context.Context, in *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	log.Printf("Received SayHello: %v", in.GetName())
	return &helloworld.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func (s *server) SayHelloAgain(ctx context.Context, in *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	log.Printf("Received SayHelloAgain: %v", in.GetName())
	return &helloworld.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func (s *server) StreamData(in *helloworld.HelloStreamRequest, stream helloworld.GreeterService_StreamDataServer) error {
	log.Printf("Received StreamData:")

	k := int32(2)
	N := in.GetNumber()
	for N > 1 {
		if N%k == 0 {
			N = N - 10
			//send to client
			stream.Send(&helloworld.HelloStreamReply{
				Message: N,
			})
			time.Sleep(1000 * time.Millisecond)
		} else {
			k++
			log.Printf("k increase to %v", k)
		}
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	helloworld.RegisterGreeterServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
