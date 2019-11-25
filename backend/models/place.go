package models

type Place struct {
	//Id   string `json:"id"`
	Name string `json:"name"`
}

/*func CreatePlace(p NewPlace) interface{} {
	res, err := PLACES.InsertOne(context.TODO(), p)
	utils.CheckErr(err)
	return res.InsertedID
}*/

func UpdatePlace(p Place) bool {
	return true
}

func DeletePlace(p Place) bool {
	return true
}

func GetPlaces(un string) []Place {
	u := GetUserByUserName(un)
	cs := u.Places
	return cs
}

func PrintPlaces(c []Place) {
	for i := range c {
		print(c[i].Name, "\n")
	}
}
