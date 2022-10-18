package main

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func connection() {
	if err := godotenv.Load(); err != nil {
		log.Println("Could not load .env file")
	}

	uri := os.Getenv("MONGO_URI")

	if uri == "" {
		log.Fatal("MONGO_URI is not set")
	}

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