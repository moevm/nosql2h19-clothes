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

type BufCloth struct {
	Name       string      `json:"name"`
	Color      string      `json:"color"`
	Notes      string      `json:"notes"`
	Img        string      `json:"img"`
	CategoryId interface{} `json:"cid"`
}

func CreateCloth(c NewCloth) interface{} {
	if CATEGORIES.FindOne(context.TODO(), c.CategoryName) != nil {
		res, err := CATEGORIES.InsertOne(context.TODO(), c)
		utils.CheckErr(err)
		cloth := BufCloth{
			Name:       c.Name,
			Color:      c.Color,
			Notes:      c.Notes,
			Img:        c.Img,
			CategoryId: res,
		}
		res_2, err_2 := CLOTHES.InsertOne(context.TODO(), cloth)
		utils.CheckErr(err_2)
		return res_2.InsertedID
	}
	return nil
}

/*
func CreateClothes(c []interface{}) interface{} {
	res, err := CLOTHES.InsertMany(context.TODO(), c)
	utils.CheckErr(err)
	return res.InsertedIDs
}
*/

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
