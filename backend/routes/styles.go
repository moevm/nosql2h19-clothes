package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"nosql2h19-clothes/backend/models"
	"nosql2h19-clothes/backend/utils"
)

var styleTitle = "Style"

func GetStyles(c *gin.Context) {
	//idParam := c.Param(utils.IdKey)
	idParam := "Ilya Bykov"
	res := models.GetStyles(idParam)
	c.JSON(http.StatusOK, res)
}

func AddStyle(c *gin.Context) {
	user := "Ilya Bykov"
	var w models.Style
	if err := c.ShouldBindJSON(&w); err == nil {
		res := models.AddStyle(user, w)
		if res == true {
			c.JSON(http.StatusOK, ApiStatusSimple{utils.SuccessMessage})
		} else {
			c.JSON(http.StatusUnprocessableEntity, ApiMessage{utils.EntityNotUpdatedMessage(styleTitle)})
		}
	} else {
		c.JSON(http.StatusBadRequest, err)
	}
}

func UpdateStyle(c *gin.Context) {

}

func DeleteStyle(c *gin.Context) {
	user := "Ilya Bykov"
	idParam := c.Param(utils.IdKey)
	w := models.Style{Name: idParam}
	res := models.DeleteStyle(user, w)
	res2 := models.DeleteClothByStyle(user, w.Name)
	if res == true && res2 == true {
		c.JSON(http.StatusOK, ApiMessage{utils.SuccessMessage})
	} else {
		c.JSON(http.StatusBadRequest, ApiMessage{utils.EntityNotDeletedMessage(styleTitle)})
	}
}
