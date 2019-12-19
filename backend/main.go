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
		fmt.Println("Database successful initialisation")
		createFirstUsers()
		/* ... */
	} else {
		log.Panic("Database not initialisation")
	}
	models.LoadNewUsers(pathToUsersDb)
	fmt.Println("\nGET TEST for Ilya Bykov:\n")
	id := models.GetUserIdByUserName("Ilya Bykov")
	fmt.Println(id)
	models.PrintCategories(models.GetCategories("Ilya Bykov"))
	models.PrintClothes(models.GetClothes("Ilya Bykov"))
	models.PrintPlaces(models.GetPlaces("Ilya Bykov"))
	models.PrintGroups(models.GetGroups("Ilya Bykov"))
	fmt.Println("\n-------------------------------------\n")
	users := models.GetUsers()
	fmt.Println(users)
	//p := models.Place{Name: "job"}
	//models.DeletePlace("Ilya Bykov", p)

	router := gin.Default()
	router.Use(CORSMiddleware())

	router.Use(CORSMiddleware())
	router.Static("/api/upload", "./api/upload")
	router.Static("/api/tmp", "./api/tmp")

	//The jwt middleware

	userAuthMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:         "secret zone",
		Key:           []byte("secret key"),
		Timeout:       24 * 720 * time.Hour,
		MaxRefresh:    24 * 720 * time.Hour,
		Authenticator: routes.AuthenticatorUser,
		Authorizator:  routes.AuthorizatorUser,
		Unauthorized:  routes.UnauthorizedUser,
		TokenLookup:   "header:Authorization",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
		IdentityKey:   identityKey,
		//PayloadFunc:     payloadFunc,
		IdentityHandler: identityHandler,
	})

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	//API
	api := router.Group("/api")
	api.POST("/login", userAuthMiddleware.LoginHandler)

	api.Use( /*userAuthMiddleware.MiddlewareFunc()*/ )
	{
		// added api groups later
		admin := api.Group("/admin")
		{
			admin.GET("/users", routes.GetUsers)
			admin.POST("", routes.UploadFile)

		}
		user := api.Group("/home") //TODO change to id
		{
			places := user.Group("/places")
			{
				places.POST("", routes.AddPlace)
				places.GET("", routes.GetPlaces)
				places.PUT("/:id", routes.UpdatePlace)
				places.DELETE("/:id", routes.DeletePlace)
			}
			styles := user.Group("/styles")
			{
				styles.POST("", routes.AddStyle)
				styles.GET("", routes.GetStyles)
				styles.PUT("/:id", routes.UpdateStyle)
				styles.DELETE("/:id", routes.DeleteStyle)

			}
			clothes := user.Group("clothes")
			{
				clothes.POST("", routes.AddCloth)
				clothes.GET("", routes.GetClothes)
				clothes.PUT("/:id", routes.UpdateCloth)
				clothes.DELETE("/:id", routes.DeleteCloth)

			}
			categories := user.Group("/categories")
			{
				categories.POST("", routes.AddCategory)
				categories.GET("", routes.GetCategories)
				categories.PUT("/:id", routes.UpdateCategory)
				categories.DELETE("/:id", routes.DeleteCategory)

			}
			groups := user.Group("/groups")
			{
				groups.POST("", routes.AddGroup)
				groups.GET("", routes.GetGroups)
				groups.PUT("/:id", routes.UpdateGroup)
				groups.DELETE("/:id", routes.DeleteGroup)

			}
		}
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

/*func payloadFunc(data interface{}) jwt.MapClaims {
	if v, ok := data.(string); ok {
		return jwt.MapClaims{
			identityKey: v,
		}
	}
	return jwt.MapClaims{}
}*/

func identityHandler(c *gin.Context) interface{} {
	claims := jwt.ExtractClaims(c)
	return claims["id"].(string)
}

func createFirstUsers() {
	/*Username   string             `json:"Username"`
	Password   string             `json:"Password"`
	Role       string             `json:"Role"`
	Name       string             `json:"Name"`
	Email      string             `json:"Email"`
	Age        int                `json:"Age"`
	Gender     int                `json:"Gender"`
	Categories []Category
	Places     []Place
	Groups     []Group
	Styles     []Style
	Clothes    []Cloth*/
	admin := models.GetUserByEmail("admin@gmail.com")
	if admin == nil {
		newAdmin := models.NewUser{}
		newAdmin.Username = "admin"
		newAdmin.Password = routes.CryptUserPassword("admin")
		newAdmin.Role = "admin"
		newAdmin.Name = "admin"
		newAdmin.Email = "admin@gmail.com"
		newAdmin.Age = 45
		newAdmin.Gender = 0
		/*newAdmin.Categories = nil
		newAdmin.Places = nil
		newAdmin.Groups = nil
		newAdmin.Styles = nil
		newAdmin.Clothes = nil*/
		res := models.CreateUser(newAdmin)
		fmt.Println(res)
	}
	Gosha := models.GetUserByEmail("gosha@gmail.com")
	if Gosha == nil {
		newGosha := models.NewUser{}
		newGosha.Username = "gosha"
		newGosha.Password = routes.CryptUserPassword("gosha")
		newGosha.Role = "user"
		newGosha.Name = "Gosha Ivanov"
		newGosha.Email = "gosha@gmail.com"
		newGosha.Age = 21
		newGosha.Gender = 0
		newGosha.Categories = nil
		newGosha.Places = nil
		newGosha.Groups = nil
		newGosha.Styles = nil
		newGosha.Clothes = nil
		res1 := models.CreateUser(newGosha)
		fmt.Println(res1)
	}
	Ivan := models.GetUserByEmail("ivan@gmail.com")
	if Ivan == nil {
		newIvan := models.NewUser{}
		newIvan.Username = "ivan"
		newIvan.Password = routes.CryptUserPassword("ivan")
		newIvan.Role = "user"
		newIvan.Name = "Ivan Ivanov"
		newIvan.Email = "ivan@gmail.com"
		newIvan.Age = 28
		newIvan.Gender = 0
		newIvan.Categories = nil
		newIvan.Places = nil
		newIvan.Groups = nil
		newIvan.Styles = nil
		newIvan.Clothes = nil
		res2 := models.CreateUser(newIvan)
		fmt.Println(res2)
	}
	Sergey := models.GetUserByEmail("sergey@gmail.com")
	if Sergey == nil {
		newSergey := models.NewUser{}

		newSergey.Password = routes.CryptUserPassword("sergey")
		newSergey.Role = "user"
		newSergey.Name = "Sergey Ivanov"
		newSergey.Email = "sergey@gmail.com"
		newSergey.Age = 18
		newSergey.Gender = 0
		newSergey.Categories = nil
		newSergey.Places = nil
		newSergey.Groups = nil
		newSergey.Styles = nil
		newSergey.Clothes = nil
		res3 := models.CreateUser(newSergey)
		fmt.Println(res3)
	}
}
