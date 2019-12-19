package routes

import (
	"fmt"
	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	"nosql2h19-clothes/backend/models"
	"nosql2h19-clothes/backend/roles"
	"nosql2h19-clothes/backend/utils"
	"strconv"
)

func GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, models.GetUsers())
}

type login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

var validate *validator.Validate

const (
	fieldErrMsg = "Field validation for '%s' failed on the '%s' tag"
)

func getValidateError(errorValidate error) ApiErrorsWithMessage {

	if _, ok := errorValidate.(*validator.InvalidValidationError); ok {
		fmt.Println(errorValidate)
		return ApiErrorsWithMessage{ApiMessage{"InvalidValidationError"}, []ApiError{}}
	}

	apiMess := ApiMessage{"Validate error"}

	var apiErrors []ApiError
	for _, err := range errorValidate.(validator.ValidationErrors) {

		apiError := ApiError{}
		apiError.Code = 1
		apiError.Field = err.Field()
		apiError.Message = fmt.Sprintf(fieldErrMsg, err.Field(), err.Tag())
		apiErrors = append(apiErrors, apiError)
	}

	return ApiErrorsWithMessage{apiMess, apiErrors}

}

/*func CreateUser(c *gin.Context) {
	var newUserBuf models.BufferUser

	validate = validator.New()

	if err := c.ShouldBindJSON(&newUserBuf); err == nil {
		fmt.Println(newUserBuf)
		newUser := models.ToNewUser(newUserBuf)
		errorValidate := validate.Struct(newUser)

		if errorValidate != nil {
			c.JSON(http.StatusBadRequest, getValidateError(errorValidate))
			return
		}

		existUser := models.GetUserByUserName(newUser.Username)
		if existUser != nil {
			c.JSON(http.StatusUnprocessableEntity, ApiMessage{"Username exist"})
		} else {
			password := CryptUserPassword(newUser.Password)
			if password == newUser.Password {
				c.JSON(http.StatusBadRequest, ApiMessage{"Server error with password"})
				return
			}
			// Set Crypt password
			newUser.Password = password
			uid := models.CreateUser(newUser)
			c.JSON(http.StatusOK, gin.H{"username": uid})
		}

	} else {
		c.JSON(http.StatusBadRequest, err)
	}
}*/

func AuthenticatorUser(c *gin.Context) (interface{}, error) {

	var loginVals login
	if err := c.ShouldBind(&loginVals); err != nil {
		return "", jwt.ErrMissingLoginValues
	}
	userId := loginVals.Username
	password := loginVals.Password

	if len(userId) > 0 && len(password) > 0 {
		existUserAuth := models.GetUserAuthByUserName(userId)
		fmt.Println(existUserAuth)
		if existUserAuth != nil {
			if IsCorrectUserPassword(existUserAuth.Password, password) {
				return userId, nil
			}
		}
	}

	return nil, jwt.ErrFailedAuthentication

}

func AuthorizatorUser(user interface{}, _ *gin.Context) bool {

	userId, _ := user.(string)
	existUserAuth := models.GetUserAuthByUserName(userId)
	if existUserAuth != nil {
		if roles.Rbac.IsGranted(existUserAuth.Role, roles.Permissions["admin"], nil) {
			return true
		}
	}

	return false
}

func UnauthorizedUser(c *gin.Context, code int, message string) {
	c.JSON(code, ApiMessage{message})
}

/*func UpdateUser(c *gin.Context) {
	var buf_user models.BufferUser

	if err := c.ShouldBindJSON(&buf_user); err == nil {
		user := models.ToNewUser(buf_user)
		if len(user.Username) > 0 {
			existUser := models.GetUserByUserName(user.Username)

			if existUser == nil {
				c.JSON(http.StatusUnprocessableEntity, gin.H{"status": "fail", "id": existUser.Username})
				return
			}

			if models.UpdateUserFull(user) {
				c.JSON(http.StatusOK, gin.H{"success": "ok"})
			} else {
				c.JSON(http.StatusBadRequest, ApiMessage{"not updated"})
			}
		} else {
			// TODO: Change Error Message
			c.JSON(400, ApiMessage{"username not valid"})
		}
	} else {
		c.AbortWithError(400, err)
	}
}*/
/*
func DeleteUser(c *gin.Context) {
	id := utils.StringToInt(c.Param("id"))
	fmt.Println(id)
	if id > 0 {
		existUser := models.GetUserById(id)
		if existUser == nil {
			c.JSON(http.StatusUnprocessableEntity, ApiMessage{"User not exist"})
		} else {
			if models.DeleteUser(*existUser) {
				c.JSON(http.StatusOK, ApiMessage{"success"})
			} else {
				c.JSON(http.StatusBadRequest, ApiMessage{"Server can not delete user"})
			}
		}
	} else {
		c.JSON(400, ApiMessage{"id not valid"})
	}
}
*/

func GetUserById(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	utils.CheckErr(err)

	if id != 0 {
		user := models.GetUserById(id)
		if user != nil {
			c.JSON(http.StatusOK, user)
		} else {
			c.JSON(http.StatusNotFound, ApiMessage{"user with id not found"})
		}
	} else {
		c.JSON(400, ApiMessage{"id not valid"})
	}
}

/*
func GetUserInfo(c *gin.Context) {

	claims := jwt.ExtractClaims(c)
	username := claims["id"].(string)
	userId := models.GetUserIdByUserName(username)
	if userId == 0 {
		c.JSON(400, ApiMessage{"username not valid"})
		return
	}

	user := models.GetUserById(userId)
	if user != nil {
		c.JSON(http.StatusOK, user)
	} else {
		c.JSON(http.StatusNotFound, ApiMessage{"User with id not found"})
	}

}
*/
func CryptUserPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		// TODO: Properly handle error
		fmt.Println(err)
		return password
	}

	fmt.Println("Hash to store:", string(hash))
	return string(hash)
}

func IsCorrectUserPassword(hash string, password string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); err != nil {
		// TODO: Properly handle error
		fmt.Println(err)
		return false
	}
	fmt.Println("Password was correct!")
	return true
}
