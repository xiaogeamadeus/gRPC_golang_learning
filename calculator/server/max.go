package main

import (
	pb "github.com/Clement-Jean/grpc-go-course/calculator/proto"
	"io"
	"log"
	"math"
)

func (s *Server) Max(stream pb.CalculatorService_MaxServer) error {
	log.Println("Max function was invoked")
	currentMax := int32(0)
	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}
		currentMax = int32(math.Max(float64(req.Num), float64(currentMax)))
		err = stream.Send(&pb.MaxResponse{
			Max: currentMax,
		})

		if err != nil {
			log.Fatalf("Error while sending data to client: %v\n", err)
		}

	}
}
