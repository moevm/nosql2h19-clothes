package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"nosql2h19-clothes/backend/models"
	"nosql2h19-clothes/backend/utils"
)

var groupTitle = "Group"

func GetGroups(c *gin.Context) {
	//idParam := c.Param(utils.IdKey)
	idParam := "Ilya Bykov"
	groups := models.GetGroups(idParam)
	c.JSON(http.StatusOK, groups)
}

func GetGroupByDate(c *gin.Context) {
	idParam := c.Param(utils.IdKey)
	user := "Ilya Bykov"
	groups := models.GetGroupByDate(user, idParam)
	c.JSON(http.StatusOK, groups)
}

func AddGroup(c *gin.Context) {
	user := "Ilya Bykov"
	var w models.Group
	if err := c.ShouldBindJSON(&w); err == nil {
		res := models.AddGroup(user, w)
		if res == true {
			c.JSON(http.StatusOK, ApiStatusSimple{utils.SuccessMessage})
		} else {
			c.JSON(http.StatusUnprocessableEntity, ApiMessage{utils.EntityNotUpdatedMessage(groupTitle)})
		}
	} else {
		c.JSON(http.StatusBadRequest, err)
	}
}

func UpdateGroup(c *gin.Context) {

}

func DeleteGroup(c *gin.Context) {
	/*user := "Ilya Bykov"
	idParam := c.Param(utils.IdKey)
	w := models.Group{Name: idParam}
	res := models.DeleteGroup(user, w)
	if res == true {
		c.JSON(http.StatusOK, ApiMessage{utils.SuccessMessage})
	} else {
		c.JSON(http.StatusBadRequest, ApiMessage{utils.EntityNotDeletedMessage(groupTitle)})
	}*/
}
