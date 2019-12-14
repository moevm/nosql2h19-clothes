package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"nosql2h19-clothes/backend/models"
	"nosql2h19-clothes/backend/utils"
)

var placeTitle = "Place"

func GetPlaces(c *gin.Context) {
	//idParam := c.Param(utils.IdKey)
	idParam := "Ilya Bykov"
	places := models.GetPlaces(idParam)
	c.JSON(http.StatusOK, places)
}

func AddPlace(c *gin.Context) {
	user := "Ilya Bykov"
	var w models.Place
	if err := c.ShouldBindJSON(&w); err == nil {
		res := models.AddPlace(user, w)
		if res == true {
			c.JSON(http.StatusOK, ApiStatusSimple{utils.SuccessMessage})
		} else {
			c.JSON(http.StatusUnprocessableEntity, ApiMessage{utils.EntityNotUpdatedMessage(placeTitle)})
		}
	} else {
		c.JSON(http.StatusBadRequest, err)
	}
}

func UpdatePlace(c *gin.Context) {

}

func DeletePlace(c *gin.Context) {
	user := "Ilya Bykov"
	idParam := c.Param(utils.IdKey)
	w := models.Place{Name: idParam}
	res := models.DeletePlace(user, w)
	if res == true {
		c.JSON(http.StatusOK, ApiMessage{utils.SuccessMessage})
	} else {
		c.JSON(http.StatusBadRequest, ApiMessage{utils.EntityNotDeletedMessage(placeTitle)})
	}
}
