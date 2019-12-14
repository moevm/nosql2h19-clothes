package models

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"nosql2h19-clothes/backend/utils"
)

type Style struct {
	//Id   string `json:"id"`
	Name string `json:"name"`
}

func GetStyles(un string) []Style {
	u := GetUserByUserName(un)
	cs := u.Styles
	return cs
}

func DeleteStyle(un string, c Style) bool {
	u := GetUserByUserName(un)
	updateResult, err := USERS.UpdateOne(context.TODO(), bson.D{{"_id", u.Id}}, bson.D{{"$pull", bson.D{{"styles", c}}}})
	utils.CheckErr(err)
	fmt.Println("update result: ", updateResult)
	return true
}

func AddStyle(un string, c Style) bool {
	u := GetUserByUserName(un)
	updateResult, err := USERS.UpdateOne(context.TODO(), bson.D{{"_id", u.Id}}, bson.D{{"$addToSet", bson.D{{"styles", c}}}})
	utils.CheckErr(err)
	fmt.Println("update result: ", updateResult)
	return true
}

func PrintStyles(c []Style) {
	for i := range c {
		print(c[i].Name, "\n")
	}
}
