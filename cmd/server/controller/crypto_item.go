package controller

import (
	"upvote/grpc/proto/pb"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CryptoItem struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	CryptoName  string             `bson:"crypto_name"`
	CryptoLikes int32              `bson:"crypto_likes"`
}

func DocumentToCrypto(data *CryptoItem) *pb.Crypto {
	return &pb.Crypto{
		Id:          data.ID.Hex(),
		CryptoName:  data.CryptoName,
		CryptoLikes: 0,
	}
}
