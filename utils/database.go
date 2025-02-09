package utils

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client // Declaring a global variable ensuring that all handlers can access the MongoDB client

func ConnectDB() *mongo.Client {
	mongoURI := os.Getenv("MONGODB_URI") // Fetching the MongoDB URI from the environment variables to keep credentials secure
	if mongoURI == "" {
		log.Fatal("MONGODB_URI environment variable not set") // Exits the application if the MongoDB URI is not set
	}

	clientOptions := options.Client().ApplyURI(mongoURI).SetServerSelectionTimeout(10 * time.Second)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second) // Configuring the MongoDB with a 10-second timeout, preventing long-waits if the database is unreachable
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions) // Connecting to the MongoDB database
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err) // Exits the application if the connection to the MongoDB database fails
	}

	err = client.Ping(ctx, nil) // Verify the connection to the MongoDB database by sending a ping
	if err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}

	log.Println("Connected to MongoDB!")
	return client
}

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("url_shortener").Collection(collectionName)
	return collection
}

func InitDB() {
	Client = ConnectDB()
}
