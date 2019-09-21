package models

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var DB *mongo.Database
var USERS *mongo.Collection
var NEWUSERS *mongo.Collection
var USERSAUTH *mongo.Collection
var DAYS *mongo.Collection
var CATEGORIES *mongo.Collection
var PLACES *mongo.Collection

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

	db := client.Database("NoSQL-clothes")
	DB = db
	return true
}

func initCollections(db *mongo.Database) {
	collectionUsers := db.Collection("Users")
	collectionNewUsers := db.Collection("NewUsers")
	collectionUsersAuth := db.Collection("UsersAuth")
	collectionDays := db.Collection("Days")
	collectionCategories := db.Collection("Category")
	collectionPlaces := db.Collection("Places")
	USERS = collectionUsers
	NEWUSERS = collectionNewUsers
	USERSAUTH = collectionUsersAuth
	DAYS = collectionDays
	CATEGORIES = collectionCategories
	PLACES = collectionPlaces
}

func Migrate(db *mongo.Database) {

}
