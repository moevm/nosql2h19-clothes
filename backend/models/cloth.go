package models

type Cloth struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	Color string `json:"color"`
	Notes string `json:"notes"`
	Img   string `json:"img"`
}
