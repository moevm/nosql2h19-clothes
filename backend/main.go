package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"nosql2h19-clothes/backend/models"
	"os"
)

var isDev = false

var isBuild = "dev"

var pathToUsersDb = "./api/tmp/db_newUser.json"
var pathToTest = "./api/tmp/test.json"

type Dev struct {
	DBPort string `json:"port"`
}
type Prod struct {
	DBPort string `json:"port"`
}
type Configuration struct {
	Dev  `json:"dev"`
	Prod `json:"prod"`
}

func main() {

	if isBuild == "dev" {
		isDev = true
		fmt.Println("Start Developing Mode")
	} else if isBuild == "prod" {
		isDev = false
		fmt.Println("Start Prodaction Mode")
	}

	if models.InitDB(ParceConfig(isDev)) {
		fmt.Println("Database successfull initialisation")
		/* ... */
	} else {
		log.Panic("Database not initialisation")
	}
	//models.LoadNewUsers(pathToUsersDb)
	id := models.GetUserIdByUserName("Ilya Bykov")
	fmt.Println("\n", id, "\n")
	models.PrintCategories(models.GetCategories("Ilya Bykov"))
	models.PrintClothes(models.GetClothes("Ilya Bykov"))
	models.PrintPlaces(models.GetPlaces("Ilya Bykov"))
	models.PrintGroups(models.GetGrops("Ilya Bykov"))

	router := gin.Default()
	router.Use(CORSMiddleware())

	router.Run(":5000")
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length,X-Requested-With, Accept-Encoding, X-CSRF-Token, Authorization")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			fmt.Println("OPTIONS")
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}

func ParceConfig(isDev bool) string {

	file, _ := os.Open("conf.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("Develop Port: " + configuration.Dev.DBPort)
	fmt.Println("Production Port: " + configuration.Prod.DBPort)
	if isDev {
		return configuration.Dev.DBPort
	}
	return configuration.Prod.DBPort
}
