package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"nosql2h19-clothes/examples/test_1/utils"
	//"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
)

type Trainer struct {
	Name string
	Age  int
	City string
}

func main() {
	fmt.Println("MongoDB EXAMPLE START")

	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	utils.CheckErr(err)

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	utils.CheckErr(err)

	fmt.Println("Connected to MongoDB!")

	// Get a handle for your collection
	collection := client.Database("test").Collection("trainers")

	// Some dummy data to add to the Database
	ilya := Trainer{"Ilya", 1, "Bugulma"}
	katya := Trainer{"Katya", 2, "Spb"}
	artem := Trainer{"Artem", 3, "Sochi"}

	// Insert a single document
	insertResult, err := collection.InsertOne(context.TODO(), ilya)
	utils.CheckErr(err)

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)

	// Insert multiple documents
	trainers := []interface{}{katya, artem}

	insertManyResult, err := collection.InsertMany(context.TODO(), trainers)
	utils.CheckErr(err)

	fmt.Println("Inserted multiple documents: ", insertManyResult.InsertedIDs)

	// Update a document
	filter := bson.D{{"name", "Ilya"}}

	update := bson.D{
		{"$inc", bson.D{
			{"age", 5},
		}},
	}

	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
	utils.CheckErr(err)

	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)

	// Find a single document
	var result Trainer

	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	utils.CheckErr(err)

	fmt.Printf("Found a single document: %+v\n", result)

	findOptions := options.Find()
	findOptions.SetLimit(2)

	var results []*Trainer

	// Finding multiple documents returns a cursor
	cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
	utils.CheckErr(err)

	// Iterate through the cursor
	for cur.Next(context.TODO()) {
		var elem Trainer
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, &elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	// Close the cursor once finished
	cur.Close(context.TODO())

	fmt.Printf("Found multiple documents (array of pointers): %+v\n", results)

	// Delete all the documents in the collection
	deleteResult, err := collection.DeleteMany(context.TODO(), bson.D{{}})
	utils.CheckErr(err)

	fmt.Printf("Deleted %v documents in the trainers collection\n", deleteResult.DeletedCount)

	// Close the connection once no longer needed
	err = client.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Connection to MongoDB closed.")
	}

}
