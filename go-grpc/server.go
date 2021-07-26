package main

import (
	"context"
	"io"
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
	log.Printf("SayHello called: %v", in.GetName())
	return &helloworld.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func (s *server) SayHelloAgain(ctx context.Context, in *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	log.Printf("SayHelloAgain called: %v", in.GetName())
	return &helloworld.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func (s *server) ServerStreamData(in *helloworld.ServerStreamRequest, stream helloworld.GreeterService_ServerStreamDataServer) error {
	log.Printf("ServerStreamData called:")

	k := int32(2)
	N := in.GetNumber()
	for N < 200 {
		if N%k == 0 {
			N = N + 10
			//send to client
			stream.Send(&helloworld.ServerStreamReply{
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

func (*server) ClientStreamData(stream helloworld.GreeterService_ClientStreamDataServer) error {
	log.Printf("ClientStreamData called:")
	var total float32

	for {
		request, error := stream.Recv()

		// khi client đã gửi xong thì trả về tổng của các số
		if error == io.EOF {
			rsp := &helloworld.ClientStreamReply{Total: total}
			time.Sleep(2 * time.Second)
			return stream.SendAndClose(rsp)
		}

		if error != nil {
			log.Fatalf("error: %v", error)
		}

		log.Printf("Number from request: %v", request.GetNumber())
		total += request.GetNumber()
	}

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
