package main

import (
	"context"
	pb "github.com/Clement-Jean/grpc-go-course/calculator/proto"
	"io"
	"log"
)

func doPrime(c pb.CalculatorServiceClient) {
	log.Printf("doPrime was invoked")

	req := &pb.PrimeRequest{
		Num: 120,
	}

	stream, err := c.Primes(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling Primes: %v\n", err)
	}

	for {
		res, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading the stream: %v\n", err)
		}

		log.Printf("Primes: %d\n", res.Res)
	}
}
