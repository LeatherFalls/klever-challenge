package main

import (
	"context"
	"log"
	"upvote/grpc/proto/pb"
)

func create(client pb.CryptoServiceClient) string {
	log.Printf("Creating a new crypto...")

	crypto := &pb.Crypto{
		CryptoName: "KLV",
	}

	res, err := client.Create(context.Background(), crypto)

	if err != nil {
		log.Fatalf("Error while creating crypto: %v", err)
	}

	log.Printf("Crypto created: %v", res.String())

	return res.Id
}

func listAll(client pb.CryptoServiceClient) {
	
}

func listById(client pb.CryptoServiceClient, id string) *pb.Crypto {
	return nil
}

func update(client pb.CryptoServiceClient, id string) {
	
}

func delete(client pb.CryptoServiceClient, id string) {

}

func upvoteCrypto(client pb.CryptoServiceClient, id string) {

}