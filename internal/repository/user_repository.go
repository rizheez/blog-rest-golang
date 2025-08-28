package repository

import (
	"blog-rest/internal/models"
	"errors"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetAllUsers() ([]models.User, error)
	GetUserById(id uint) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	CreateUser(user *models.User) error
	UpdateUser(user *models.User) error
	DeleteUser(user *models.User) error
}
type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}
func (r *userRepository) GetAllUsers() ([]models.User, error) {
	var users []models.User
	result := r.db.Find(&users)
	err := result.Error
	return users, err
}

func (r *userRepository) GetUserById(id uint) (*models.User, error) {
	var user models.User
	if err := r.db.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) CreateUser(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) UpdateUser(user *models.User) error {
	return r.db.Model(&models.User{}).Where("id = ?", user.ID).Updates(user).Error
}

func (r *userRepository) DeleteUser(user *models.User) error {
	return r.db.Delete(user).Error
}
