package models

type Category struct {
	Id      int64   `json:"id"`
	Name    string  `json:"name"`
	Clothes []Cloth `json:"clothes"`
}

func CreateCategory(c Category) int64 {
	var uid int64

	return uid
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

func GetCategoryById(id int64) Category {
	var c Category
	return c
}

func GetCategoryByName(name string) Category {
	var c Category
	return c
}
