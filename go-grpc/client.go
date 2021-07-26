package main

import (
	"context"
	"fmt"
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
	ServerStreamData(client)
	//ClientStreamData(client)
}

func UnaryGrpc(client helloworld.GreeterServiceClient) {
	r, err := client.SayHello(context.Background(), &helloworld.HelloRequest{Name: defaultName})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}

func ServerStreamData(client helloworld.GreeterServiceClient) {
	stream, err3 := client.ServerStreamData(context.Background(), &helloworld.ServerStreamRequest{Number: 0})
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

func ClientStreamData(client helloworld.GreeterServiceClient) {
	// ví dụ:
	// gửi các số dương liên tục lên server và server sẽ trả về tổng của các số
	stream, err3 := client.ClientStreamData(context.Background())
	if err3 != nil {
		log.Fatalf("error: %v", err3)
	}

	listReq := []helloworld.ClientStreamRequest{}

	for i := 0; i < 10; i++ {
		listReq = append(listReq, helloworld.ClientStreamRequest{Number: float32(i)})
	}

	for _, req := range listReq {
		errorSend := stream.Send(&req)
		if errorSend != nil {
			log.Fatalf("error send: %v", err3)
		}
	}

	response, errorCloseAndRecv := stream.CloseAndRecv()
	if errorCloseAndRecv != nil {
		log.Fatalf("error CloseAndRecv: %v", errorCloseAndRecv)
	}

	fmt.Println("Total: ", response.GetTotal())

}
