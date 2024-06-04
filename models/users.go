package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"hello/packages/database"
	u "hello/utils"
)

type User struct {
	gorm.Model
	Username string `gorm:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

func (user *User) Validate() (map[string]interface{}, bool) {

	if user.Email == "" {
		return u.Message(false, "Email address is required"), false
	}

	if len(user.Password) < 6 {
		return u.Message(false, "Password is required"), false
	}

	return u.Message(false, "passed"), true
}

func (user *User) Create() map[string]interface{} {
	if resp, ok := user.Validate(); !ok {
		return resp
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	database.GetDB().Create(user)

	user.Password = ""

	response := u.Message(true, "Account has been created")
	response["user"] = user
	return response
}

func Login(email, password string) map[string]interface{} {

	user := &User{}

	err := database.GetDB().Table("accounts").Where("email = ?", email).First(user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return u.Message(false, "Email not found")
		}
		return u.Message(false, "500")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil && errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		return u.Message(false, "Invalid login credentials")
	}

	response := u.Message(true, "Logged In")
	response["user"] = user

	return response
}
