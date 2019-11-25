package models

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"nosql2h19-clothes/examples/test_1/utils"
)

type UserAuth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type NewUser struct {
	Username   string `json:"username" validate:"required,gte=3"`
	Password   string `json:"password" validate:"required,gte=6"`
	Name       string `json:"name" validate:"required,gte=3"`
	Email      string `json:"email" validate:"required,email"`
	Age        int64  `json:"age"`
	Gender     int64  `json:"gender"`
	Role       int    `json:"Role"`
	Categories []Category
	Places     []Place
	Groups     []Group
	Styles     []Style
	Clothes    []Cloth
}

type User struct {
	Id         primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Username   string             `json:"Username"`
	Password   string             `json:"Password"`
	Role       int                `json:"Role"`
	Name       string             `json:"Name"`
	Email      string             `json:"Email"`
	Age        int                `json:"Age"`
	Gender     int                `json:"Gender"`
	Categories []Category
	Places     []Place
	Groups     []Group
	Styles     []Style
	Clothes    []Cloth
}

func GetUserByUserName(username string) *User {
	var result User
	err := USERS.FindOne(context.TODO(), bson.D{{"name", username}}).Decode(&result)
	utils.CheckErr(err)
	return &result
}

func GetUserAuthByUserName(username string) *UserAuth {
	user := UserAuth{}

	return &user
}

func GetUserIdByUserName(username string) interface{} {
	var result User
	err := USERS.FindOne(context.TODO(), bson.D{{"name", username}}).Decode(&result)
	utils.CheckErr(err)
	return result.Id
}

func CreateUser(u NewUser) interface{} {
	insertResult, err := USERS.InsertOne(context.TODO(), u)
	utils.CheckErr(err)
	var result User
	err = USERS.FindOne(context.TODO(), insertResult).Decode(&result)
	utils.CheckErr(err)
	return result.Id
}

func UpdateUser(userAuth NewUser) bool {

	return true
}

func DeleteUser(user User) bool {

	return true
}

func GetUsers() []User {
	var users []User
	return users
}

func GetUserById(id int64) *User {

	user := User{}

	return &user
}

func GetUserByUsername(username string) *User {

	user := User{}

	return &user
}
