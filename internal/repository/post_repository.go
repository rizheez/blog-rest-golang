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

func GetPostById(id int) (models.Post, error) {
	post := models.Post{}
	result := database.DB.Preload("User").Preload("Category").First(&post, id)
	err := result.Error
	return post, err
}

func CreatePost(post *models.Post) error {
	return database.DB.Create(post).Error
}

func UpdatePost(post *models.Post) error {
	return database.DB.Model(&models.Post{}).Where("id = ?", post.ID).Updates(post).Error

}

func DeletePost(post *models.Post) error {
	return database.DB.Delete(post).Error
}
