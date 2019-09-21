package models

import (
	"context"
	"nosql2h19-clothes/backend/utils"
)

type Cloth struct {
	Id         int64  `json:"id"`
	Name       string `json:"name"`
	Color      string `json:"color"`
	Notes      string `json:"notes"`
	Img        string `json:"img"`
	CategoryId int64  `json:"cid"`
}

type NewCloth struct {
	Name         string `json:"name"`
	Color        string `json:"color"`
	Notes        string `json:"notes"`
	Img          string `json:"img"`
	CategoryName string `json:"cname"`
}

func CreateCloth(c NewCloth) interface{} {
	res, err := CLOTHES.InsertOne(context.TODO(), c)
	utils.CheckErr(err)
	return res.InsertedID
}

func CreateClothes(c []interface{}) interface{} {
	res, err := CLOTHES.InsertMany(context.TODO(), c)
	utils.CheckErr(err)
	return res.InsertedIDs
}

func UpdateCloth(c Cloth) bool {
	return true
}

func DeleteCloth(c Cloth) bool {
	return true
}

func GetClothes() []Cloth {
	var cs []Cloth
	return cs
}

func GetClothById(id int64) *Cloth {
	var c Cloth

	return &c
}

func GetClothesByName(name string) []Cloth {
	var cs []Cloth
	return cs
}
