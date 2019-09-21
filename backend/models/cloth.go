package models

type Cloth struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	Color string `json:"color"`
	Notes string `json:"notes"`
	Img   string `json:"img"`
}

func CreateCloth(c Cloth) int64 {
	var uid int64

	return uid
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

func GetClothById(id int64) Cloth {
	var c Cloth
	return c
}

func GetClothesByName(name string) []Cloth {
	var cs []Cloth
	return cs
}
