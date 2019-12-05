package models

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"nosql2h19-clothes/backend/utils"
)

type Category struct {
	//Id   string `json:"id"`
	Name string `json:"name"`
}

/*func CreateCategory(c NewCategory) interface{} {
	res, err := CATEGORIES.InsertOne(context.TODO(), c)
	utils.CheckErr(err)
	return res.InsertedID
}*/

func UpdateCategory(c Category) bool {
	return true
}

func DeleteCategory(un string, c Category) bool {
	u := GetUserByUserName(un)
	updateResult, err := USERS.UpdateOne(context.TODO(), bson.D{{"_id", u.Id}}, bson.D{{"$pull", bson.D{{"categories", c}}}})
	utils.CheckErr(err)
	fmt.Println("update result: ", updateResult)
	return true
}

func GetCategories(un string) []Category {
	u := GetUserByUserName(un)
	cs := u.Categories
	return cs
}

func PrintCategories(c []Category) {
	for i := range c {
		print(c[i].Name, "\n")
	}
}

func AddCategory(un string, c Category) bool {
	u := GetUserByUserName(un)
	updateResult, err := USERS.UpdateOne(context.TODO(), bson.D{{"_id", u.Id}}, bson.D{{"$addToSet", bson.D{{"categories", c}}}})
	utils.CheckErr(err)
	fmt.Println("update result: ", updateResult)
	return true
}

func GetCategoryById(id int64) *Category {
	var c Category
	return &c
}

func GetCategoriesByName(name string) []Category {
	var c []Category
	return c
}
