package services

import (
	"blog-rest/internal/models"
	"blog-rest/internal/repository"
)

func GetAllCategories() ([]models.Category, error) {
	return repository.GetAllCategories()
}

func GetCategoriesByid(id int) (models.Category, error) {
	return repository.GetCategoriesByid(id)
}

func CreateCategory(category *models.Category) error {
	return repository.CreateCategory(category)
}

func UpdateCategory(category *models.Category) error {
	return repository.UpdateCategory(category)
}

func DeleteCategory(category *models.Category) error {
	return repository.DeleteCategory(category)
}
