package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Ctx = context.Background()

// ConnectMongoDB establishes a connection to MongoDB and returns the client and collection
func ConnectMongoDB() (*mongo.Client, *mongo.Collection) {
	ctx, cancel := context.WithTimeout(Ctx, 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://mongo:27017"))
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}

	collection := client.Database("userdb").Collection("users")
	return client, collection
}
