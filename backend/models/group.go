package models

type Group struct {
	//Id         string `json:"id"`
	Clothes    []Cloth
	Date       string `json:"date"`
	Place_name string `json:"place_name"`
}
