package models

import "github.com/jinzhu/gorm"

type Post struct {
	gorm.Model
	UserId        *int   `json:"user_id,omitempty"`
	Title         string `json:"title"`
	Content       string `json:"content"`
	CommentsCount int    `json:"comments_count"`
}
