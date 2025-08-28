package services

import (
	"blog-rest/internal/dto"
	"blog-rest/internal/models"
	"blog-rest/internal/repository"
	"blog-rest/internal/utils"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Register(dto.RegisterDTO) (*models.User, error)
	Login(dto.LoginDTO) (map[string]string, error)
}

func NewAuthService(u repository.UserRepository) AuthService {
	return &authService{userRepo: u}
}

type authService struct {
	userRepo repository.UserRepository
}

func (s *authService) Register(input dto.RegisterDTO) (*models.User, error) {
	// check existing
	exist, err := s.userRepo.GetUserByEmail(input.Email)
	if err != nil {
		return nil, err
	}
	if exist != nil {
		return nil, errors.New("email already used")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: string(hashed),
	}

	if err := s.userRepo.CreateUser(user); err != nil {
		return nil, err
	}
	// clear password before returning
	user.Password = ""
	return user, nil
}

func (s *authService) Login(input dto.LoginDTO) (map[string]string, error) {
	user, err := s.userRepo.GetUserByEmail(input.Email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return nil, errors.New("invalid credentials")
	}

	token, err := utils.GenerateToken(user.ID, time.Hour*24) // token 24 jam
	if err != nil {
		return nil, err
	}

	return token, nil
}
