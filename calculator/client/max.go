package main

import (
	"context"
	pb "github.com/Clement-Jean/grpc-go-course/calculator/proto"
	"io"
	"log"
	"time"
)

func doMax(c pb.CalculatorServiceClient) {
	log.Println("doMax function was invoked")

	stream, err := c.Max(context.Background())

	if err != nil {
		log.Fatalf("Error while creating stream: %v\n", err)
	}

	reqs := []*pb.MaxRequest{
		{Num: 2},
		{Num: 2},
		{Num: 1231312},
		{Num: 5},
		{Num: 10000},
	}

	waitc := make(chan struct{})

	go func() {
		for _, req := range reqs {
			log.Printf("Sending request: %v\n", req)
			stream.Send(req)
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Printf("Error while receiving: %v\n", err)
				break
			}
			log.Printf("Received: %v\n", res.Max)

		}
		close(waitc)
	}()

	<-waitc
}
