package controller

import (
	"context"
	"log"
	"net"
	"upvote/grpc/proto/pb"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
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
	request = &pb.Crypto{
		CryptoName:  "Bitcoin",
		CryptoLikes: 100,
	}

	insert, _ := Collection.InsertOne(context.Background(), request)

	res := &pb.CryptoId{
		Id: insert.InsertedID.(primitive.ObjectID).Hex(),
	}
	return res, nil
}

func (s *CryptoServiceServer) ListAllCrypto(_ *emptypb.Empty, stream pb.CryptoService_ListAllServer) error {
	for {
		streamRes := &pb.Crypto{
			Id:          "634f6df1f56fbac87d062fbe",
			CryptoName:  "Bitcoin",
			CryptoLikes: 100,
		}

		stream.Send(streamRes)
	}
}

func (s *CryptoServiceServer) ListCryptoById(ctx context.Context, request *pb.CryptoId) (*pb.Crypto, error) {
	request = &pb.CryptoId{
		Id: "634f6df1f56fbac87d062fbe",
	}

	res := &pb.Crypto{
		Id:          "634f6df1f56fbac87d062fbe",
		CryptoName:  "Bitcoin",
		CryptoLikes: 100,
	}
	return res, nil
}

func (s *CryptoServiceServer) UpdateCrypto(ctx context.Context, request *pb.Crypto) (*pb.Crypto, error) {
	request = &pb.Crypto{
		Id:          "634f6df1f56fbac87d062fbe",
		CryptoName:  "Bitcoin",
		CryptoLikes: 100,
	}

	cryptoUpdated := &pb.Crypto{
		Id:          "634f6df1f56fbac87d062fbe",
		CryptoName:  "BitcoinUpdated",
		CryptoLikes: 100,
	}

	Collection.UpdateOne(
		context.Background(),
		bson.M{"_id": request.Id},
		bson.M{"$set": cryptoUpdated},
	)

	return cryptoUpdated, nil
}

func (s *CryptoServiceServer) DeleteCrypto(ctx context.Context, request *pb.CryptoId) (*pb.CryptoId, error) {
	crypto := &pb.Crypto{
		CryptoName:  "Bitcoin",
		CryptoLikes: 100,
	}

	newCrypto, _ := s.CreateCrypto(ctx, crypto)

	log.Printf("aaaaaaaaaa %v", newCrypto)

	deleted, _ := Collection.DeleteOne(context.Background(), newCrypto.Id)

	log.Print(deleted)

	res := &pb.CryptoId{
		Id: crypto.Id,
	}

	return res, nil
}
