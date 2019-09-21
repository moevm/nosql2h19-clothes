package models

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var DB *mongo.Database

func InitDB(port string) bool {

	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:" + port)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
		return false
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
		return false
	}

	db := client.Database("test")

	DB = db

	return true
}
