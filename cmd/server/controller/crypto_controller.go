package controller

import (
	"context"
	"fmt"
	"log"
	"upvote/grpc/proto/pb"

	"go.mongodb.org/mongo-driver/bson"
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
	log.Println("ListAll was invoked")
	
	ctx := context.Background()

	cur, err := Collection.Find(ctx, primitive.D{{}})

	if err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal server error: %v", err),
		)
	}

	defer cur.Close(ctx)

	for cur.Next(ctx) {
		data := &CryptoItem{}
		
		err := cur.Decode(data)
		
		if err != nil {
			return status.Errorf(
				codes.Internal,
				fmt.Sprintf("Error while decoding data from MongoDB: %v", err),
			)
		}
		
		stream.Send(DocumentToCrypto(data))
	}
	
	if err = cur.Err(); err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unknown internal error: %v", err),
		)
	}
	
	return nil
}

func (*CryptoServer) ListById(ctx context.Context, in *pb.CryptoId) (*pb.Crypto, error) {
	log.Printf("ListById was invoked with %v\n", in)

	oid, err := primitive.ObjectIDFromHex(in.Id)

	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"Cannot parse ID",
		)
	}

	data := &CryptoItem{}

	log.Print(data)
	
	filter := bson.M{"_id": oid}

	res := Collection.FindOne(ctx, filter)

	log.Print(res)

	if err := res.Decode(data); err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Cannot find crypto with specified ID: %v", err),
		)
	}

	return DocumentToCrypto(data), nil
}

func (*CryptoServer) Update(ctx context.Context, in *pb.Crypto) (*emptypb.Empty, error) {
	log.Printf("UpdateCrypto was invoked with %v\n", in)

	oid, err := primitive.ObjectIDFromHex(in.Id)

	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"Cannot parse ID",
		)
	}

	data := &CryptoItem{
		CryptoName: in.CryptoName,
	}

	log.Print(data)

	res, err := Collection.UpdateOne(
		ctx,
		bson.M{"_id": oid},
		bson.M{"$set": data},
	)

	log.Print(res)

	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			"Could not update",
		)
	}

	if res.MatchedCount == 0 {
		return nil, status.Errorf(
			codes.NotFound,
			"Cannot find crypto with ID",
		)
	}

	return &emptypb.Empty{}, nil
}

func (*CryptoServer) Delete(ctx context.Context, in *pb.CryptoId) (*emptypb.Empty, error) {
	oid, err := primitive.ObjectIDFromHex(in.Id)

	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"Cannot parse OID",
		)
	}

	res, err := Collection.DeleteOne(ctx, bson.M{ "_id": oid })

	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Cannot delete object: %v", err),
		)
	}

	if res.DeletedCount == 0 {
		return nil, status.Errorf(
			codes.NotFound,
			"Crypto not found",
		)
	}

	return &emptypb.Empty{}, nil
}

func (*CryptoServer) UpvoteCrypto(ctx context.Context, in *pb.CryptoId) (*emptypb.Empty, error) {
	return nil, nil
}