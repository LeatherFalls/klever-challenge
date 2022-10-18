package main

import (
	"log"
	"net"
	"upvote/grpc/proto/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	connection()

	lis, err := net.Listen("tcp", "localhost:50051")

	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}

	log.Print("Server listening on localhost:50051")

	s := grpc.NewServer()
	pb.RegisterCryptoServiceServer(s, &Server{})

	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}