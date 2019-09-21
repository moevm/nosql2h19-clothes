package models

type Place struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

func CreatePlace(p Place) int64 {
	var uid int64

	return uid
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

func GetPlaceById(id int64) Place {
	var p Place
	return p
}

func GetPlacesByName(name string) []Place {
	var ps []Place
	return ps
}
