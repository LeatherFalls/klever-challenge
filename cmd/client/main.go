package main

import (
	"log"
	"upvote/grpc/proto/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var address = "localhost:50051"

func main() {
	connection, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Error while connecting to server: %v", err)
	}

	defer connection.Close()

	client := pb.NewCryptoServiceClient(connection)
	
	id := create(client)
	listAll(client)
	listById(client, id)
	update(client, id)
	delete(client, id)
	upvoteCrypto(client, id)
}