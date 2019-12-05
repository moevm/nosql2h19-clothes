package models

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"nosql2h19-clothes/backend/utils"
)

type Place struct {
	//Id   string `json:"id"`
	Name string `json:"name"`
}

/*func CreatePlace(p NewPlace) interface{} {
	res, err := PLACES.InsertOne(context.TODO(), p)
	utils.CheckErr(err)
	return res.InsertedID
}*/

func UpdatePlace(p Place) bool {
	return true
}

func DeletePlace(un string, c Place) bool {
	u := GetUserByUserName(un)
	updateResult, err := USERS.UpdateOne(context.TODO(), bson.D{{"_id", u.Id}}, bson.D{{"$pull", bson.D{{"places", c}}}})
	utils.CheckErr(err)
	fmt.Println("update result: ", updateResult)
	return true
}

func GetPlaces(un string) []Place {
	u := GetUserByUserName(un)
	cs := u.Places
	return cs
}

func AddPlace(un string, c Place) bool {
	u := GetUserByUserName(un)
	updateResult, err := USERS.UpdateOne(context.TODO(), bson.D{{"_id", u.Id}}, bson.D{{"$addToSet", bson.D{{"places", c}}}})
	utils.CheckErr(err)
	fmt.Println("update result: ", updateResult)
	return true
}

func PrintPlaces(c []Place) {
	for i := range c {
		print(c[i].Name, "\n")
	}
}
