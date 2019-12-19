package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
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
}

func UnloadFile(c *gin.Context) {

}
