package controller

import (
	"context"
	"log"
	"net"
	"upvote/grpc/proto/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type CryptoServiceServer struct {
	pb.UnimplementedCryptoServiceServer
}

func StartServer() {
	lis, err := net.Listen("tcp", "localhost:50051")
	
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	server := grpc.NewServer()
	pb.RegisterCryptoServiceServer(server, &CryptoServer{})

	log.Printf("Server listening at: %v", lis.Addr())

	reflection.Register(server)

	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func (s *CryptoServiceServer) CreateCrypto(ctx context.Context, request *pb.Crypto) (*pb.CryptoId, error) {
	res := &pb.CryptoId{
		Id: "634f6df1f56fbac87d062fbe",
	}
	return res, nil
}