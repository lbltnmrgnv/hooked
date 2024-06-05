package handlers

import (
	"encoding/json"
	"errors"
	"github.com/jinzhu/gorm"
	"hello/internal/models"
	"hello/packages/database"
	res "hello/packages/http"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		res.Response(w, res.ErrorResponse(4000, "Invalid request"), http.StatusBadRequest)
		return
	}
	if err := user.Validate(); err != nil {
		res.Response(w, err, http.StatusBadRequest)
		return
	}
	createdUser := user.Create()
	res.Response(w, map[string]interface{}{"user": createdUser}, http.StatusOK)
}

func Login(w http.ResponseWriter, r *http.Request) {
	userReq := &models.LoginRequest{}
	err := json.NewDecoder(r.Body).Decode(userReq)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		res.Response(w, res.ErrorResponse(4000, "Invalid request"), http.StatusBadRequest)
		return
	}

	user := database.Postgres().Where("username = ?", userReq.Username).First(&models.User{})
	if errors.Is(user.Error, gorm.ErrRecordNotFound) {
		res.Response(w, res.ErrorResponse(4010, "Invalid username or password"), http.StatusUnauthorized)
		return
	}

}
