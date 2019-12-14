package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"nosql2h19-clothes/backend/models"
)

func GetClothes(c *gin.Context) {
	//idParam := c.Param(utils.IdKey)
	idParam := "Ilya Bykov"
	res := models.GetClothes(idParam)
	c.JSON(http.StatusOK, res)
}
