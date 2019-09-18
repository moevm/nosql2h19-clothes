package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"nosql2h19-clothes/examples/test_1/utils"
)

func main() {

	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	utils.CheckErr(err)

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	utils.CheckErr(err)

	fmt.Println("Connected to MongoDB!")

}
