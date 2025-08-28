package repository

import (
	"blog-rest/internal/models"
	"errors"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	GetAllCategories() ([]models.Category, error)
	GetCategoriesByid(id uint) (*models.Category, error)
	CreateCategory(category *models.Category) error
	UpdateCategory(category *models.Category) error
	DeleteCategory(category *models.Category) error
}
type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db}
}
func (r *categoryRepository) GetAllCategories() ([]models.Category, error) {
	var category []models.Category

	result := r.db.Find(&category)
	err := result.Error

	return category, err
}

func (r *categoryRepository) GetCategoriesByid(id uint) (*models.Category, error) {
	var category models.Category
	result := r.db.First(&category, id)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &category, nil
}

func (r *categoryRepository) CreateCategory(category *models.Category) error {
	return r.db.Create(category).Error
}

func (r *categoryRepository) UpdateCategory(category *models.Category) error {
	return r.db.Model(&models.Category{}).Where("id = ?", category.ID).Updates(category).Error
}

func (r *categoryRepository) DeleteCategory(category *models.Category) error {
	return r.db.Delete(category).Error
}
