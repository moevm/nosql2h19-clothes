package models

import "C"
import (
	"context"
	"nosql2h19-clothes/backend/utils"
)

type Category struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type NewCategory struct {
	Name string `json:"name"`
}

func CreateCategory(c NewCategory) interface{} {
	res, err := CATEGORIES.InsertOne(context.TODO(), c)
	utils.CheckErr(err)
	return res.InsertedID
}

func UpdateCategory(c Category) bool {
	return true
}

func DeleteCategory(c Category) bool {
	return true
}

func GetCategories() []Category {
	var cs []Category

	return cs
}

func GetCategoryById(id int64) *Category {
	var c Category
	return &c
}

func GetCategoriesByName(name string) []Category {
	var c []Category
	return c
}
