package models

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"io/ioutil"
	"nosql2h19-clothes/backend/utils"
	"path/filepath"
	"reflect"
	"time"
)

var DB *mongo.Database
var USERS *mongo.Collection
var CTX context.Context

//var NEWUSERS *mongo.Collection
//var USERSAUTH *mongo.Collection
//var DAYS *mongo.Collection
//var CATEGORIES *mongo.Collection
//var PLACES *mongo.Collection
//var CLOTHES *mongo.Collection
var Users []User

func InitDB(port string) bool {

	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:" + port)
	fmt.Println("clientOptions TYPE:", reflect.TypeOf(clientOptions))
	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	utils.CheckErr(err)

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	utils.CheckErr(err)

	db := client.Database("NoSQL-clothes")
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)
	DB = db
	print(DB.Name(), "\n")
	CTX = ctx
	initCollections(DB)
	return true
}

func initCollections(db *mongo.Database) {
	collectionUsers := db.Collection("Users")
	//collectionNewUsers := db.Collection("NewUsers")
	//collectionUsersAuth := db.Collection("UsersAuth")
	//collectionDays := db.Collection("Days")
	//collectionCategories := db.Collection("Category")
	//collectionPlaces := db.Collection("Places")
	//collectionClothes := db.Collection("Clothes")
	USERS = collectionUsers
	fmt.Println("Collection type:", reflect.TypeOf(collectionUsers))
	//NEWUSERS = collectionNewUsers
	//USERSAUTH = collectionUsersAuth
	//DAYS = collectionDays
	//CATEGORIES = collectionCategories
	//PLACES = collectionPlaces
	//CLOTHES = collectionClothes
}

func Migrate(db *mongo.Database) {

}

func LoadNewUsers(path string) {
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)
	CTX = ctx
	docPath, err := filepath.Abs(path)
	utils.CheckErr(err)
	byteValues, err := ioutil.ReadFile(docPath)
	utils.CheckErr(err)
	fmt.Println("ioutil.ReadFile byteValues TYPE:", reflect.TypeOf(byteValues))
	//fmt.Println("byteValues:", byteValues, "n")
	//fmt.Println("byteValues:", string(byteValues))
	err = json.Unmarshal(byteValues, &Users)
	utils.CheckErr(err)
	fmt.Println("MongoFields Users:", reflect.TypeOf(Users))
	for i := range Users {
		print(Users[i].Name, "\n")
	}
	print("\n")
	for i := range Users {
		user := Users[i]
		result, insertErr := USERS.InsertOne(ctx, user)
		utils.CheckErr(insertErr)
		print("InsertOne() API result:", result)
	}
}
