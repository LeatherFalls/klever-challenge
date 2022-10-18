package main

import (
	"context"
	"io"
	"log"
	"upvote/grpc/proto/pb"

	"google.golang.org/protobuf/types/known/emptypb"
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
	log.Printf("Listing all crypto...")

	list, err := client.ListAll(context.Background(), &emptypb.Empty{})

	if err != nil {
		log.Fatalf("Error while reading crypto: %v", err)
	}

	for {
		response, err := list.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading stream: %v", err)
		}

		log.Printf("Crypto: %v", response)
	}
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