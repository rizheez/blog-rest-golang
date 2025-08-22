package services

import (
	"blog-rest/internal/models"
	"blog-rest/internal/repository"
)

func GetPosts() ([]models.Post, error) {
	return repository.GetAllPosts()

}

func GetPostById(id int) (models.Post, error) {
	return repository.GetPostById(id)
}

func CreatePost(post *models.Post) error {
	return repository.CreatePost(post)
}

func UpdatePost(post *models.Post) error {
	return repository.UpdatePost(post)
}

func DeletePost(post *models.Post) error {
	return repository.DeletePost(post)
}
