package handlers

import (
	"encoding/json"
	"hello/internal/models"
	"hello/internal/service"
	res "hello/packages/http"
	"net/http"
	"strconv"
)

func CreatePost() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		post := &models.Post{}
		err := json.NewDecoder(r.Body).Decode(post)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			res.Response(w, res.Error(4000, "Invalid request"), http.StatusBadRequest)
			return
		}

		userIdStr := r.Context().Value("uid").(string)
		userId, err := strconv.Atoi(userIdStr)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		post.UserId = &userId
		createdPost := service.CreatePost(post)
		createdPost.UserId = nil
		res.Response(w, map[string]interface{}{"post": createdPost}, http.StatusOK)
	}
}

func ListOfPosts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		userIdStr := r.Context().Value("uid").(string)
		userId, err := strconv.Atoi(userIdStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		posts := service.ListPostsByUserId(userId)
		for i := range posts {
			posts[i].UserId = nil
		}

		res.Response(w, map[string]interface{}{"posts": posts}, http.StatusOK)
	}
}
