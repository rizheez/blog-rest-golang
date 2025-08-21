package repository

import (
	"blog-rest/internal/database"
	"blog-rest/internal/models"
)

func GetAllPosts() ([]models.Post, error) {
	var posts []models.Post
	result := database.DB.Preload("User").Preload("Category").Find(&posts)
	err := result.Error
	return posts, err
}

func CreatePost(post *models.Post) error {
	return database.DB.Create(post).Error
}
