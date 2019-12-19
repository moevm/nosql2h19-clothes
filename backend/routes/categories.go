package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"nosql2h19-clothes/backend/models"
	"nosql2h19-clothes/backend/utils"
)

var categoryTitle = "Category"

func GetCategories(c *gin.Context) {
	//idParam := c.Param(utils.IdKey)
	idParam := "Ilya Bykov"
	res := models.GetCategories(idParam)
	c.JSON(http.StatusOK, res)
}

func AddCategory(c *gin.Context) {
	user := "Ilya Bykov"
	var w models.Category
	if err := c.ShouldBindJSON(&w); err == nil {
		res := models.AddCategory(user, w)
		if res == true {
			c.JSON(http.StatusOK, ApiStatusSimple{utils.SuccessMessage})
		} else {
			c.JSON(http.StatusUnprocessableEntity, ApiMessage{utils.EntityNotUpdatedMessage(categoryTitle)})
		}
	} else {
		c.JSON(http.StatusBadRequest, err)
	}
}

func UpdateCategory(c *gin.Context) {

}

func DeleteCategory(c *gin.Context) {
	user := "Ilya Bykov"
	idParam := c.Param(utils.IdKey)
	w := models.Category{Name: idParam}
	res := models.DeleteCategory(user, w)
	res2 := models.DeleteClothByCategory(user, w.Name)
	if res == true && res2 == true {
		c.JSON(http.StatusOK, ApiMessage{utils.SuccessMessage})
	} else {
		c.JSON(http.StatusBadRequest, ApiMessage{utils.EntityNotDeletedMessage(categoryTitle)})
	}
}
