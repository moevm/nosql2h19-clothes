package models

type Category struct {
	Id      int64   `json:"id"`
	Name    string  `json:"name"`
	Clothes []Cloth `json:"clothes"`
}
