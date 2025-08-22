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

func GetCategoriesByid(id int) (models.Category, error) {
	var category models.Category
	result := database.DB.First(&category, id)
	err := result.Error
	return category, err
}

func CreateCategory(category *models.Category) error {
	return database.DB.Create(category).Error
}

func UpdateCategory(category *models.Category) error {
	return database.DB.Model(&models.Category{}).Where("id = ?", category.ID).Updates(category).Error
}

func DeleteCategory(category *models.Category) error {
	return database.DB.Delete(category).Error
}
