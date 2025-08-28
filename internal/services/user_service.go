package services

import (
	"blog-rest/internal/models"
	"blog-rest/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	GetAllUsers() ([]models.User, error)
	GetUserById(id uint) (*models.User, error)
	CreateUser(user *models.User) error
	UpdateUser(user *models.User) error
	DeleteUser(user *models.User) error
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(u repository.UserRepository) UserService {
	return &userService{userRepo: u}
}

func (s *userService) GetAllUsers() ([]models.User, error) {
	return s.userRepo.GetAllUsers()
}

func (s *userService) GetUserById(id uint) (*models.User, error) {
	return s.userRepo.GetUserById(id)
}

func (s *userService) CreateUser(user *models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return s.userRepo.CreateUser(user)
}

func (s *userService) UpdateUser(user *models.User) error {
	return s.userRepo.UpdateUser(user)
}

func (s *userService) DeleteUser(user *models.User) error {
	return s.userRepo.DeleteUser(user)
}
