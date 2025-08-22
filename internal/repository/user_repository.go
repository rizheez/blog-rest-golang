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

func GetUserById(id int) (models.User, error) {
	var user models.User
	result := database.DB.First(&user, id)
	err := result.Error
	return user, err
}

func CreateUser(user *models.User) error {
	return database.DB.Create(user).Error
}

func UpdateUser(user *models.User) error {
	return database.DB.Model(&models.User{}).Where("id = ?", user.ID).Updates(user).Error
}

func DeleteUser(user *models.User) error {
	return database.DB.Delete(user).Error
}
