package controller

import (
	"context"
	"log"
	"testing"
	"upvote/grpc/proto/pb"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	go StartServer()
}

func connection() {
	uri := "mongodb+srv://root:nsyPGGrHvSzklZCC@cluster0.y9tu6ly.mongodb.net/?retryWrites=true&w=majority"

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))

	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(context.Background())

	if err != nil {
		log.Fatal(err)
	}

	Collection = client.Database("cryptodb").Collection("crypto")
}

func TestCreate(t *testing.T) {
	s := CryptoServiceServer{}
	connection()

	crypto := &pb.Crypto{
		CryptoName:  "Bitcoin",
		CryptoLikes: 100,
	}

	res, err := s.CreateCrypto(context.Background(), crypto)

	if err != nil {
		t.Errorf("Error while creating crypto: %v", err)
	}

	crypto.Id = res.Id

	log.Print(res.Id)
	log.Print(crypto.Id)

	if res.Id != crypto.Id {
		t.Errorf("Invalid response from server: %v", res)
	}

	if res.Id == crypto.Id {
		log.Print("Crypto created successfully")
	}

	defer Collection.Drop(context.Background())
}

func TestListById(t *testing.T) {
	s := CryptoServiceServer{}
	connection()

	request := &pb.CryptoId{
		Id: "634f6df1f56fbac87d062fbe",
	}

	res, err := s.ListCryptoById(context.Background(), request)

	if err != nil {
		t.Errorf("Error while listing crypto by id: %v", err)
	}

	if res.Id != request.Id {
		t.Errorf("Invalid response from server: %v", res)
	}
}

func TestUpdate(t *testing.T) {
	s := CryptoServiceServer{}
	connection()

	request := &pb.Crypto{
		Id:          "634f6df1f56fbac87d062fbe",
		CryptoName:  "Bitcoin",
		CryptoLikes: 100,
	}

	res, err := s.UpdateCrypto(context.Background(), request)

	if res.CryptoName == request.CryptoName {
		t.Errorf("CryptoName must to be updated: %v", err)
	}

	if err != nil {
		t.Errorf("Error while updating crypto: %v", err)
	}

	defer Collection.Drop(context.Background())
}

func TestDelete(t *testing.T) {
	s := CryptoServiceServer{}
	connection()
	defer Collection.Drop(context.Background())

	request := &pb.CryptoId{
		Id: "634f6df1f56fbac87d062fbe",
	}

	res, err := s.DeleteCrypto(context.Background(), request)

	if res.Id == "" {
		log.Print("Crypto deleted")
	}

	if err != nil {
		t.Errorf("Error while deleting crypto: %v", err)
	}
}

func TestUpvoteCrypto(t *testing.T) {
	s := CryptoServiceServer{}
	connection()

	c := &pb.Crypto{
		CryptoName:  "Bitcoin",
		CryptoLikes: 100,
	}

	test, _ := s.CreateCrypto(context.Background(), c)

	oid, _ := primitive.ObjectIDFromHex(test.Id)

	Collection.UpdateOne(
		context.Background(),
		bson.M{"_id": oid},
		bson.M{"$inc": bson.M{"cryptolikes": 1}},
	)

	crypto, err := Collection.FindOne(context.Background(), bson.M{"_id": oid}).DecodeBytes()

	e, _ := crypto.Elements()

	cl := e[3].Value().Int32()

	if cl != 101 {
		t.Errorf("Error while upvoting crypto: %v", err)
	}

	if err != nil {
		t.Errorf("Error while upvoting crypto: %v", err)
	}

	if cl == 101 {
		log.Print("Crypto upvoted successfully")
	}

	defer Collection.Drop(context.Background())
}
