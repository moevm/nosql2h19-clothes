package main

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"nosql2h19-clothes/examples/test_1/utils"
	"os"

	//"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
)

var isDev = false

var isBuild = "dev"

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

	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	utils.CheckErr(err)

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	utils.CheckErr(err)

	fmt.Println("Connected to MongoDB!")

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
	fmt.Println(configuration.Dev.DBPort)
	fmt.Println(configuration.Prod.DBPort)
	if isDev {
		return configuration.Dev.DBPort
	}
	return configuration.Prod.DBPort
}
