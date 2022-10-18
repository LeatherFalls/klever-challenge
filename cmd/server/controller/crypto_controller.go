package controller

import (
	"context"
	"fmt"
	"log"
	"upvote/grpc/proto/pb"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type CryptoServer struct {
	pb.CryptoServiceServer
}

func NewCryptoController() CryptoServer {
	return CryptoServer{}
}

func (*CryptoServer) Create(ctx context.Context, in *pb.Crypto) (*pb.CryptoId, error) {
	log.Printf("CreateCrypto was invoked with %v\n", in)

	data := CryptoItem{
		CryptoName:  in.CryptoName,
		CryptoLikes: 0,
	}

	res, err := Collection.InsertOne(ctx, data)

	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal server error: %v", err),
		)
	}

	oid, ok := res.InsertedID.(primitive.ObjectID)

	if !ok {
		return nil, status.Errorf(
			codes.Internal,
			"Cannot convert to OID",
		)
	}

	return &pb.CryptoId{
		Id: oid.Hex(),
	}, nil
}

func (s *CryptoServer) ListAll(_ *emptypb.Empty, stream pb.CryptoService_ListAllServer) error {
	return nil
}

func (*CryptoServer) ListById(ctx context.Context, in *pb.CryptoId) (*pb.Crypto, error) {
	return nil, nil
}

func (*CryptoServer) Update(ctx context.Context, in *pb.Crypto) (*emptypb.Empty, error) {
	return nil, nil
}

func (*CryptoServer) Delete(ctx context.Context, in *pb.CryptoId) (*emptypb.Empty, error) {
	return nil, nil
}

func (*CryptoServer) UpvoteCrypto(ctx context.Context, in *pb.CryptoId) (*emptypb.Empty, error) {
	return nil, nil
}