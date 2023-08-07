package main

import (
	pb "github.com/Clement-Jean/grpc-go-course/calculator/proto"
	"io"
	"log"
)

func (s *Server) Avg(stream pb.CalculatorService_AvgServer) error {
	log.Println("Avg function was invoked")
	sum := 0
	count := 0

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.AvgResponse{
				Result: float64(sum) / float64(count),
			})
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}

		sum += int(req.Num)
		count++
		log.Printf("sum: %f, count: %f, res: %f\n", float64(sum), float64(count), float64(sum)/float64(count))
	}
	return nil
}
