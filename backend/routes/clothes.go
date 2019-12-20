package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"nosql2h19-clothes/backend/models"
	"nosql2h19-clothes/backend/utils"
)

var clothTitle = "Cloth"

func GetClothes(c *gin.Context) {
	//idParam := c.Param(utils.IdKey)
	idParam := "Ilya Bykov"
	res := models.GetClothes(idParam)
	c.JSON(http.StatusOK, res)
}

func AddCloth(c *gin.Context) {
	user := "Ilya Bykov"
	var w models.Cloth
	if err := c.ShouldBindJSON(&w); err == nil {
		res := models.AddCloth(user, w)
		if res == true {
			c.JSON(http.StatusOK, ApiStatusSimple{utils.SuccessMessage})
		} else {
			c.JSON(http.StatusUnprocessableEntity, ApiMessage{utils.EntityNotUpdatedMessage(clothTitle)})
		}
	} else {
		c.JSON(http.StatusBadRequest, err)
	}
}

func UpdateCloth(c *gin.Context) {

}

func DeleteCloth(c *gin.Context) {
	user := "Ilya Bykov"
	idParam := c.Param(utils.IdKey)
	//w := models.Cloth{Name: idParam}
	res := models.DeleteCloth(user, idParam)
	if res == true {
		c.JSON(http.StatusOK, ApiMessage{utils.SuccessMessage})
	} else {
		c.JSON(http.StatusBadRequest, ApiMessage{utils.EntityNotDeletedMessage(clothTitle)})
	}
}
