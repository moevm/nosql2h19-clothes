package models

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"nosql2h19-clothes/examples/test_1/utils"
)

type UserAuth struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"Role"`
}

type NewUser struct {
	Username   string `json:"username" validate:"required,gte=3"`
	Password   string `json:"password" validate:"required,gte=6"`
	Name       string `json:"name" validate:"required,gte=3"`
	Email      string `json:"email" validate:"required,email"`
	Age        int64  `json:"age"`
	Gender     int64  `json:"gender"`
	Role       string `json:"Role"`
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
	Role       string             `json:"Role"`
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

type UserInfo struct {
	Id       primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Username string             `json:"Username"`
	Password string             `json:"Password"`
	Role     string             `json:"Role"`
	Name     string             `json:"Name"`
	Email    string             `json:"Email"`
	Age      int                `json:"Age"`
	Gender   int                `json:"Gender"`
}

func GetUserByUserName(username string) *User {
	var result User
	err := USERS.FindOne(context.TODO(), bson.D{{"name", username}}).Decode(&result)
	utils.CheckErr(err)
	return &result
}

func GetUserByEmail(email string) *User {
	var result User
	err := USERS.FindOne(context.TODO(), bson.D{{"email", email}}).Decode(&result)
	if err != nil {
		return nil
	}
	return &result
}

func GetUserIDByUserNameAndPassword(username string, password string) primitive.ObjectID {
	var result User
	err := USERS.FindOne(context.TODO(), bson.D{{"name", username}, {"password", password}}).Decode(&result)
	utils.CheckErr(err)
	return result.Id
}

func GetUserAuthByUserName(username string) *UserAuth {
	user := UserAuth{}
	var result User
	err := USERS.FindOne(context.TODO(), bson.D{{"name", username}}).Decode(&result)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	user.Username = result.Username
	user.Password = result.Password
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
	fmt.Println(insertResult)
	utils.CheckErr(err)
	var result User
	err = USERS.FindOne(context.TODO(), bson.D{{"email", u.Email}}).Decode(&result)
	utils.CheckErr(err)
	return result.Id
}

func UpdateUser(userAuth User) bool {

	return true
}

func GetUsers() []User {
	var result []User
	cur, err := USERS.Find(context.TODO(), bson.D{{}})
	for cur.Next(context.TODO()) {
		var elem User
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		result = append(result, elem)
	}
	utils.CheckErr(err)

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.TODO())

	return result
}

func GetUserById(id int64) *User {

	user := User{}

	return &user
}

func GetUserByUsername(username string) *User {

	user := User{}

	return &user
}

func GetUsersEmails() []string {
	var result []User
	cur, err := USERS.Find(context.TODO(), bson.D{{}})
	for cur.Next(context.TODO()) {
		var elem User
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		result = append(result, elem)
	}
	if err != nil {
		//log.Fatal(err)
		return nil
	}
	if err := cur.Err(); err != nil {
		//log.Fatal(err)
		return nil
	}
	var emails []string
	for _, c := range result {
		emails = append(emails, c.Email)
	}
	cur.Close(context.TODO())
	return emails
}

func DeleteUser(email string) bool {
	//updateResult, err := USERS.UpdateOne(context.TODO(), bson.D{{"email", email}}, bson.D{{"$remove", bson.D{{"email", email}}}})
	updateResult, err := USERS.DeleteOne(context.TODO(), bson.D{{"email", email}})
	utils.CheckErr(err)
	fmt.Println("update result: ", updateResult)
	return true
}
