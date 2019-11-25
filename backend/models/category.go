package models

type Category struct {
	Name string `json:"name"`
}

/*func CreateCategory(c NewCategory) interface{} {
	res, err := CATEGORIES.InsertOne(context.TODO(), c)
	utils.CheckErr(err)
	return res.InsertedID
}*/

func UpdateCategory(c Category) bool {
	return true
}

func DeleteCategory(c Category) bool {
	return true
}

func GetCategories(u string) []Category {

}
