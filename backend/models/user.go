package models

type UserAuth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type NewUser struct {
	Username string `json:"username" validate:"required,gte=3"`
	Password string `json:"password" validate:"required,gte=6"`
	Name     string `json:"name" validate:"required,gte=3"`
	Email    string `json:"email" validate:"required,email"`
}

type User struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Email    string `json:"email"`
}

func GetUserByUserName(username string) *User {
	user := User{}

	return &user
}

func GetUserAuthByUserName(username string) *UserAuth {
	user := UserAuth{}

	return &user
}

func GetUserIdByUserName(username string) int64 {

	var userId int64

	return userId
}

func CreateUser(u NewUser) string {

	var uid string

	return uid
}

func UpdateUser(userAuth NewUser) bool {

	return true
}

func DeleteUser(user User) bool {

	return true
}

func GetUsers() []User {
	var users []User
	return users
}

func GetUserById(id int64) *User {

	user := User{}

	return &user
}

func GetUserByUsername(username string) *User {

	user := User{}

	return &user
}
