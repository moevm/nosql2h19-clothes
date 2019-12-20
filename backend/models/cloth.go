package models

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"nosql2h19-clothes/examples/test_1/utils"
)

type Cloth struct {
	//Id            string `json:"id"`
	Name         string `json:"name"`
	Color        string `json:"color"`
	Notes        string `json:"notes"`
	Img          string `json:"img"`
	CategoryName string `json:"categoryName"`
	StyleName    string `json:"styleName"`
}

/*func CreateCloth(c NewCloth) interface{} {
	if CATEGORIES.FindOne(context.TODO(), c.Category_name) != nil {
		res, err := CATEGORIES.InsertOne(context.TODO(), c)
		utils.CheckErr(err)
		return res.InsertedID
	}
	return nil
}*/

/*
func CreateClothes(c []interface{}) interface{} {
	res, err := CLOTHES.InsertMany(context.TODO(), c)
	utils.CheckErr(err)
	return res.InsertedIDs
}
*/

func DeleteCloth(un string, cl string) bool {
	u := GetUserByUserName(un)
	c := GetClothByName(un, cl)
	fmt.Println(c)
	updateResult, err := USERS.UpdateOne(context.TODO(), bson.D{{"_id", u.Id}}, bson.D{{"$pull", bson.D{{"clothes", c}}}})
	utils.CheckErr(err)
	fmt.Println("update result: ", updateResult)
	return true
}

func DeleteClothByStyle(un string, s string) bool {
	u := GetUserByUserName(un)
	for _, c := range u.Clothes {
		if c.StyleName == s {
			updateResult, err := USERS.UpdateOne(context.TODO(), bson.D{{"_id", u.Id}}, bson.D{{"$pull", bson.D{{"clothes", c}}}})
			utils.CheckErr(err)
			fmt.Println("update result: ", updateResult)
		}
	}
	return true
}

func DeleteClothByCategory(un string, s string) bool {
	u := GetUserByUserName(un)
	for _, c := range u.Clothes {
		if c.CategoryName == s {
			fmt.Println("res:", c)
			updateResult, err := USERS.UpdateOne(context.TODO(), bson.D{{"_id", u.Id}}, bson.D{{"$pull", bson.D{{"clothes", c}}}})
			utils.CheckErr(err)
			fmt.Println("update result: ", updateResult)
		}
	}
	return true
}

func GetClothes(un string) []Cloth {
	u := GetUserByUserName(un)
	cs := u.Clothes
	return cs
}

func AddCloth(un string, c Cloth) bool {
	u := GetUserByUserName(un)
	updateResult, err := USERS.UpdateOne(context.TODO(), bson.D{{"_id", u.Id}}, bson.D{{"$addToSet", bson.D{{"clothes", c}}}})
	utils.CheckErr(err)
	fmt.Println("update result: ", updateResult)
	return true
}

func PrintClothes(c []Cloth) {
	for i := range c {
		print(c[i].Name, "\n")
	}
}

func GetClothByName(un string, s string) Cloth {
	u := GetUserByUserName(un)
	cs := u.Clothes
	var res Cloth
	for _, i := range cs {
		if i.Name == s {
			res = i
		}
	}
	return res
}
