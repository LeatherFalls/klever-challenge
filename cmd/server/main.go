package main

import (
	"log"
	"net"
	"upvote/grpc/proto/pb"
	"upvote/grpc/cmd/server/controller"
	
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var cryptoController controller.CryptoServer = controller.NewCryptoController()

func main() {
	connection()

	lis, err := net.Listen("tcp", "localhost:50051")

	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}

	log.Print("Server listening on localhost:50051")

	s := grpc.NewServer()
	pb.RegisterCryptoServiceServer(s, &controller.CryptoServer{})

	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}