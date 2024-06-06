package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"hello/packages/database"
	res "hello/packages/http"
	"strings"
)

type User struct {
	gorm.Model
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) ComparePasswords(password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return err
	}
	return nil
}

func (u *User) SanitizePassword() {
	u.Password = ""
}

func (u *User) UserUniqCheck() bool {
	if result := database.Postgres().Where("email = ?", u.Email).First(&u); !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return false
	}

	if result := database.Postgres().Where("username = ?", u.Username).First(&u); !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return false
	}
	return true
}

func (u *User) Validate() map[string]interface{} {
	if u.Email == "" {
		return res.Error(4000, "Email address is required")
	}
	if !strings.Contains(u.Email, "@") {
		return res.Error(4001, "Email is not valid email")
	}
	if u.Password == "" {
		return res.Error(4000, "Password is required")
	}
	if len(u.Password) < 8 {
		return res.Error(4002, "Password must be longer than 8 characters")
	}
	if !u.UserUniqCheck() {
		return res.Error(4003, "User already exists")
	}
	return nil
}
