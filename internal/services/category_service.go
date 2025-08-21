package services

import (
	"blog-rest/internal/models"
	"blog-rest/internal/repository"
)

func GetAllCategories() ([]models.Category, error) {
	return repository.GetAllCategories()
}

func CreateCategory(user *models.Category) error {
	return repository.CreateCategory(user)
}
