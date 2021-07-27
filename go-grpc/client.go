package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

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
	//ServerStreamData(client)
	//ClientStreamData(client)
	BidirectionalStream(client)
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

func BidirectionalStream(client helloworld.GreeterServiceClient) {
	// ví dụ:
	// gửi các số dương liên tục lên server
	// và server sẽ trả về liên tục số lớn nhất trong các số gửi lên
	stream, err3 := client.BidirectionalStream(context.Background())
	if err3 != nil {
		log.Fatalf("error: %v", err3)
	}

	listReq := []helloworld.BidirectionalStreamRequest{
		helloworld.BidirectionalStreamRequest{
			Number: 20,
		},
		helloworld.BidirectionalStreamRequest{
			Number: 30,
		},
		helloworld.BidirectionalStreamRequest{
			Number: 10,
		},
		helloworld.BidirectionalStreamRequest{
			Number: 100,
		},
		helloworld.BidirectionalStreamRequest{
			Number: 50,
		},
		helloworld.BidirectionalStreamRequest{
			Number: 50,
		},
		helloworld.BidirectionalStreamRequest{
			Number: 300,
		},
	}

	go func() {
		for _, req := range listReq {
			fmt.Println("Sending number:", req.Number)
			errorSend := stream.Send(&req)
			if errorSend != nil {
				log.Fatalf("error send: %v", err3)
				break
			}
			time.Sleep(1000 * time.Millisecond)
		}
		stream.CloseSend()
	}()

	c := make(chan float32)
	go func() {
		for {
			response, errorRecv := stream.Recv()
			if errorRecv == io.EOF {
				log.Println("Ending stream Recv from server ...")
				break
			}
			if errorRecv != nil {
				log.Fatalf("error errorRecv: %v", errorRecv)
				break
			}
			max := response.GetMax()
			fmt.Println("Max: ", max)
		}

		close(c)
	}()

	<-c
}
