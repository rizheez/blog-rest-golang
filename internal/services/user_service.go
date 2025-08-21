package services

import (
	"blog-rest/internal/models"
	"blog-rest/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

func GetAllUsers() ([]models.User, error) {
	return repository.GetAllUsers()
}

func CreateUser(user *models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return repository.CreateUser(user)
}
