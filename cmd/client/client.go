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
	res, err := client.ListById(context.Background(), &pb.CryptoId{Id: id})

	if err != nil {
		log.Fatalf("Error while reading crypto: %v", err)
	}

	log.Printf("Reading crypto with id: %s", id)

	return res
}

func update(client pb.CryptoServiceClient, id string) {
	log.Printf("Updating crypto with id: %s", id)

		crypto := &pb.Crypto{
			Id:      id,
			CryptoName: "KLV is the best crypto",
		}

		_, err := client.Update(context.Background(), crypto)

		if err != nil {
			log.Fatalf("Error while updating crypto: %v", err)
		}

		log.Println("Crypto was updated!")
}

func delete(client pb.CryptoServiceClient, id string) {
	log.Printf("Deleting crypto with id: %s", id)

	res, err := client.Delete(context.Background(), &pb.CryptoId{Id: id})

	if err != nil {
		log.Fatalf("Error while deleting crypto: %v", err)
	}

	log.Printf("Crypto deleted: %v", res.String())
}

func upvoteCrypto(client pb.CryptoServiceClient, id string) {

}