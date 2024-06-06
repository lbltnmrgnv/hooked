package service

import (
	"fmt"
	"hello/internal/models"
	"hello/packages/database"
)

func CreatePost(post *models.Post) *models.Post {
	database.Postgres().Create(post)
	return post
}

func ListPostsByUserId(userId int) []models.Post {
	var posts []models.Post
	result := database.Postgres().Find(&posts).Where("user_id = ?", userId).Limit(20)
	if result.Error != nil {
		fmt.Print(result.Error)
	}
	return posts
}
