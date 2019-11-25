package models

type Cloth struct {
	//Id            string `json:"id"`
	Name          string `json:"name"`
	Color         string `json:"color"`
	Notes         string `json:"notes"`
	Img           string `json:"img"`
	Category_name string `json:"category_name"`
}

/*func CreateCloth(c NewCloth) interface{} {
	if CATEGORIES.FindOne(context.TODO(), c.Category_name) != nil {
		res, err := CATEGORIES.InsertOne(context.TODO(), c)
		utils.CheckErr(err)
		return res.InsertedID
	}
	return nil
}*/

/*
func CreateClothes(c []interface{}) interface{} {
	res, err := CLOTHES.InsertMany(context.TODO(), c)
	utils.CheckErr(err)
	return res.InsertedIDs
}
*/

func UpdateCloth(c Cloth) bool {
	return true
}

func DeleteCloth(c Cloth) bool {
	return true
}

func GetClothes(un string) []Cloth {
	u := GetUserByUserName(un)
	cs := u.Clothes
	return cs
}

func PrintClothes(c []Cloth) {
	for i := range c {
		print(c[i].Name, "\n")
	}
}
