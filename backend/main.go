package main

import (
	"encoding/json"
	"fmt"
	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"log"
	"nosql2h19-clothes/backend/models"
	"nosql2h19-clothes/backend/routes"
	"os"
	"time"
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

var identityKey = "id"

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
	fmt.Println("\nGET TEST for Ilya Bykov:\n")
	id := models.GetUserIdByUserName("Ilya Bykov")
	fmt.Println(id)
	models.PrintCategories(models.GetCategories("Ilya Bykov"))
	models.PrintClothes(models.GetClothes("Ilya Bykov"))
	models.PrintPlaces(models.GetPlaces("Ilya Bykov"))
	models.PrintGroups(models.GetGroups("Ilya Bykov"))
	fmt.Println("\n-------------------------------------\n")
	p := models.Place{Name: "job"}
	models.DeletePlace("Ilya Bykov", p)

	router := gin.Default()
	router.Use(CORSMiddleware())

	router.Use(CORSMiddleware())
	router.Static("/api/upload", "./api/upload")
	router.Static("/api/tmp", "./api/tmp")

	//The jwt middleware

	userAuthMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:           "secret zone",
		Key:             []byte("secret key"),
		Timeout:         24 * 720 * time.Hour,
		MaxRefresh:      24 * 720 * time.Hour,
		Authenticator:   routes.AuthenticatorUser,
		Authorizator:    routes.AuthorizatorUser,
		Unauthorized:    routes.UnauthorizedUser,
		TokenLookup:     "header:Authorization",
		TokenHeadName:   "Bearer",
		TimeFunc:        time.Now,
		IdentityKey:     identityKey,
		PayloadFunc:     payloadFunc,
		IdentityHandler: identityHandler,
	})

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	//API
	api := router.Group("/api")
	api.POST("/login", userAuthMiddleware.LoginHandler)

	api.Use(userAuthMiddleware.MiddlewareFunc())
	{
		/*	api.GET("users/info", routes.GetUserInfo)
			api.GET("users/id/:id", routes.GetUserById)
			api.GET("statistics/manager", routes.GetTasksStatisticsForManagers)
			api.GET("statistics/managersNames", routes.GetManagers)
			managers := api.Group("/managers")
			{
				managers.POST("", routes.CreateUser)
				managers.GET("/id/:id", routes.GetUserById)
				managers.PUT("/:id", routes.UpdateUser)
				managers.DELETE("/:id", routes.DeleteUser)
				managers.GET("/managersNames", routes.GetUsers)
			}
			status := api.Group("/taskstatus")
			{
				status.POST("", routes.CreateTaskStatus)
				status.GET("/id/:id", routes.ReadTaskStatusById)
				status.PUT("/:id", routes.UpdateTaskStatus)
				status.DELETE("/:id", routes.DeleteTaskStatus)
				status.GET("", routes.GetTaskStatuses)
			}*/

	}
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

func payloadFunc(data interface{}) jwt.MapClaims {
	if v, ok := data.(string); ok {
		return jwt.MapClaims{
			identityKey: v,
		}
	}
	return jwt.MapClaims{}
}

func identityHandler(c *gin.Context) interface{} {
	claims := jwt.ExtractClaims(c)
	return claims["id"].(string)
}
