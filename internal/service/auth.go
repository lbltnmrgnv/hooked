package service

import (
	"errors"
	"github.com/jinzhu/gorm"
	"hello/internal/lib/jwt"
	"hello/internal/models"
	"hello/packages/database"
	res "hello/packages/http"
)

func Register(user *models.User) (*models.User, map[string]interface{}) {
	err := user.HashPassword()
	if err != nil {
		return nil, res.Error(5000, "System error")
	}
	if errMsg := user.Validate(); errMsg != nil {
		return nil, errMsg
	}
	database.Postgres().Create(user)
	user.SanitizePassword()
	return user, nil
}

func Login(userReq *models.LoginRequest) (string, map[string]interface{}) {
	var user models.User
	if err := database.Postgres().Where("username = ?", userReq.Username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", res.Error(4010, "Invalid username or password")
		}
	}

	if err := user.ComparePasswords(userReq.Password); err != nil {
		return "", res.Error(4010, "Invalid username or password")
	}

	token, err := jwt.NewToken(&user)
	if err != nil {
		return "", res.Error(4010, "Invalid username or password")
	}

	return token, nil
}
