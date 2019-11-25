package models

type Cloth struct {
	//Id            string `json:"id"`
	Name          string `json:"name"`
	Color         string `json:"color"`
	Notes         string `json:"notes"`
	Img           string `json:"img"`
	Category_name string `json:"category_name"`
}

type NewCloth struct {
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

func GetClothes() []Cloth {
	var cs []Cloth
	return cs
}

func GetClothById(id int64) *Cloth {
	var c Cloth

	return &c
}

func GetClothesByName(name string) []Cloth {
	var cs []Cloth
	return cs
}
