package models

import (
	"context"
	"nosql2h19-clothes/backend/utils"
)

type Place struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type NewPlace struct {
	Name string `json:"name"`
}

func CreatePlace(p NewPlace) interface{} {
	res, err := PLACES.InsertOne(context.TODO(), p)
	utils.CheckErr(err)
	return res.InsertedID
}

func UpdatePlace(p Place) bool {
	return true
}

func DeletePlace(p Place) bool {
	return true
}

func GetPlaces() []Place {
	var ps []Place
	return ps
}

func GetPlaceById(id int64) *Place {
	var p Place
	return &p
}

func GetPlacesByName(name string) []Place {
	var ps []Place
	return ps
}
