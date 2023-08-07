package main

import (
	pb "github.com/Clement-Jean/grpc-go-course/calculator/proto"
	"log"
)

func (s *Server) Primes(in *pb.PrimeRequest, stream pb.CalculatorService_PrimesServer) error {
	log.Printf("Prime function was invoked with %v\n", in)
	divisor := int64(2)
	number := in.Num
	for number > 1 {
		if number%divisor == 0 {
			stream.Send(&pb.PrimeResponse{
				Res: divisor,
			})

			number /= divisor
		} else {
			divisor++
		}
	}
	return nil
}
