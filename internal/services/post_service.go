package services

import (
	"blog-rest/internal/models"
	"blog-rest/internal/repository"
)

func GetPosts() ([]models.Post, error) {
	return repository.GetAllPosts()

}

func CreatePost(post *models.Post) error {
	return repository.CreatePost(post)
}
