package models

type Style struct {
	//Id   string `json:"id"`
	Name string `json:"name"`
}

func GetStyles(un string) []Style {
	u := GetUserByUserName(un)
	cs := u.Styles
	return cs
}
