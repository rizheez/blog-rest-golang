package repository

import (
	"blog-rest/internal/database"
	"blog-rest/internal/models"
)

func GetAllCategories() ([]models.Category, error) {
	var category []models.Category

	result := database.DB.Find(&category)
	err := result.Error

	return category, err
}

func CreateCategory(category *models.Category) error {
	return database.DB.Create(category).Error
}
