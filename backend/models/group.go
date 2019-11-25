package models

type Group struct {
	//Id         string `json:"id"`
	Clothes    []Cloth
	Date       string `json:"date"`
	Place_name string `json:"place_name"`
}

func GetGrops(un string) []Group {
	u := GetUserByUserName(un)
	cs := u.Groups
	return cs
}

func PrintGroups(c []Group) {
	for i := range c {
		print(c[i].Date, " ", c[i].Place_name, "\n")
	}
}
