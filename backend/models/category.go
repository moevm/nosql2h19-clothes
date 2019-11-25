package models

type Category struct {
	//Id   string `json:"id"`
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

func GetCategories(un string) []Category {
	u := GetUserByUserName(un)
	cs := u.Categories
	return cs
}

func PrintCategories(c []Category) {
	for i := range c {
		print(c[i].Name, "\n")
	}
}

func GetCategoryById(id int64) *Category {
	var c Category
	return &c
}

func GetCategoriesByName(name string) []Category {
	var c []Category
	return c
}
