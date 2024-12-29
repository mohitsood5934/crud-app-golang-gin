package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

// Initialize the MongoDB client and connect to the database
func InitMongoDB() (*mongo.Client, error) {
	uri := "mongodb://admin:admin123@127.0.0.1:27017" // Use appropriate URI

	// Create a context with timeout for connecting
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	log.Println("Connecting to MongoDB...")
	var err error
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
	}

	// Retry logic to wait for MongoDB to be ready
	for i := 0; i < 5; i++ {
		err = client.Ping(ctx, nil)
		if err == nil {
			log.Println("Successfully connected to MongoDB!")
			return client, nil
		}

		log.Printf("MongoDB not ready yet. Retrying... (%d/5)\n", i+1)
		time.Sleep(3 * time.Second)
	}

	return nil, fmt.Errorf("failed to ping MongoDB after multiple attempts: %w", err)
}

// SetClient assigns the MongoDB client to the package-level variable
func SetClient(c *mongo.Client) {
	log.Println("Setting MongoDB client in package-level variable")
	client = c
}

// GetCollection retrieves a MongoDB collection by name
func GetCollection(collectionName string) *mongo.Collection {
	if client == nil {
		log.Fatal("MongoDB client is not initialized. Call InitMongoDB first.")
	}
	log.Printf("Retrieving collection: %s\n", collectionName)
	db := client.Database("mydb")
	return db.Collection(collectionName)
}
