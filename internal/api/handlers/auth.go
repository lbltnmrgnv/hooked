package handlers

import (
	"encoding/json"
	"hello/internal/models"
	"hello/internal/service"
	res "hello/packages/http"
	"net/http"
)

func Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := &models.User{}
		err := json.NewDecoder(r.Body).Decode(user)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			res.Response(w, res.Error(4000, "Invalid request"), http.StatusBadRequest)
			return
		}
		createdUser, resErr := service.Register(user)
		if resErr != nil {
			res.Response(w, resErr, http.StatusBadRequest)
			return
		}
		res.Response(w, map[string]interface{}{"user": createdUser}, http.StatusCreated)
	}
}

func Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userReq := &models.LoginRequest{}
		err := json.NewDecoder(r.Body).Decode(userReq)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			res.Response(w, res.Error(4000, "Invalid request"), http.StatusBadRequest)
			return
		}

		token, responseErr := service.Login(userReq)
		if responseErr != nil {
			res.Response(w, responseErr, http.StatusBadRequest)
			return
		}

		res.Response(w, map[string]interface{}{"token": token}, http.StatusOK)
	}
}
