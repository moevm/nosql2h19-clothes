package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"nosql2h19-clothes/backend/models"
)

func GetPlaces(c *gin.Context) {
	//idParam := c.Param(utils.IdKey)
	idParam := "Ilya Bykov"
	places := models.GetPlaces(idParam)
	c.JSON(http.StatusOK, places)
}
