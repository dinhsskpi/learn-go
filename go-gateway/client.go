package main

import (
	"context"
	"log"

	demo "go-gateway-protos/demo"

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
	client := demo.NewDemoServiceClient(conn)
	CallSayHello(client)
}

func CallSayHello(client demo.DemoServiceClient) {
	r, err := client.SayHello(context.Background(), &demo.HelloRequest{Name: defaultName})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}
