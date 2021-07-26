package main

import (
	"context"
	"io"
	"log"

	helloworld "protos/helloworld"

	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "dinhpv"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()
	client := helloworld.NewGreeterServiceClient(conn)
	// UnaryGrpc(client)
	ResponseStreamData(client)
}

func UnaryGrpc(client helloworld.GreeterServiceClient) {
	r, err := client.SayHello(context.Background(), &helloworld.HelloRequest{Name: defaultName})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}

func ResponseStreamData(client helloworld.GreeterServiceClient) {
	stream, err3 := client.ResponseStreamData(context.Background(), &helloworld.ResponseStreamRequest{Number: 0})
	if err3 != nil {
		log.Fatalf("could not greet: %v", err3)
	}

	for {
		response, reciverError := stream.Recv()

		if reciverError == io.EOF {
			log.Println("Server finish streaming")
			return
		}

		if reciverError != nil {
			log.Fatalf("StreamData reciverError %v", reciverError)
		}

		log.Println("Data stream:", response.GetMessage())
	}
}
