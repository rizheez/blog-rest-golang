package services

import (
	"blog-rest/internal/models"
	"blog-rest/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

func GetAllUsers() ([]models.User, error) {
	return repository.GetAllUsers()
}

func GetUserById(id int) (models.User, error) {
	return repository.GetUserById(id)
}

func CreateUser(user *models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return repository.CreateUser(user)
}

func UpdateUser(user *models.User) error {
	return repository.UpdateUser(user)
}

func DeleteUser(user *models.User) error {
	return repository.DeleteUser(user)
}
