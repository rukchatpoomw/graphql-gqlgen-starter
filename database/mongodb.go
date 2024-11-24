package database

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client
var MongoDB *mongo.Database

func ConnectDB() {
	uri := os.Getenv("MONGODB_URI")
	dbName := os.Getenv("MONGODB_DATABASE")

	if uri == "" {
		log.Fatal("MONGODB_URI is not set in environment variables")
	}
	if dbName == "" {
		log.Fatal("MONGODB_DATABASE is not set in environment variables")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Ping the database
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Connected to MongoDB at %s", uri)
	
	MongoClient = client
	MongoDB = client.Database(dbName)
}

func GetCollection(collectionName string) *mongo.Collection {
	return MongoDB.Collection(collectionName)
}
