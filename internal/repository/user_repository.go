package repository

import (
	"blog-rest/internal/database"
	"blog-rest/internal/models"
)

func GetAllUsers() ([]models.User, error) {
	var users []models.User
	result := database.DB.Find(&users)
	err := result.Error
	return users, err
}

func CreateUser(user *models.User) error {
	return database.DB.Create(user).Error
}
