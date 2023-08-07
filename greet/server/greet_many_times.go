package main

import (
	"fmt"
	pb "github.com/Clement-Jean/grpc-go-course/greet/proto"
	"log"
)

func (s *Server) GreetManyTimes(in *pb.GreetRequest, stream pb.GreetService_GreetManyTimesServer) error {
	log.Printf("GreetManyTimes function was invoked with: %v\n", in)

	for i := 0; i < 10; i++ {
		res := fmt.Sprintf("Hello %s, number %d", in.FirstName, i)

		err := stream.Send(&pb.GreetResponse{
			Result: res,
		})

		if err != nil {
			log.Printf("Error sending GreetResponse: %v\n", err)
			return err
		}
	}

	return nil
}
