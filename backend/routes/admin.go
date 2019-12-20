package routes

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"nosql2h19-clothes/backend/models"
	"time"
)

func UploadFile(c *gin.Context) {
	file, err := c.FormFile("myFile")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}

	FileName := "api/upload/" + "File_" + time.Now().Format("20060102150405") + "_" + file.Filename
	fmt.Println(FileName)
	if err := c.SaveUploadedFile(file, FileName); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
		return
	}
	models.LoadNewUsers(FileName)
}

func UnloadFile(c *gin.Context) {
	data := models.GetUsers()
	file, _ := json.MarshalIndent(data, "", " ")
	filename := "mydb_" + time.Now().Format("20060102150405") + ".json"
	_ = ioutil.WriteFile("api/tmp/"+filename, file, 0644)
	c.JSON(http.StatusOK, filename)
}
