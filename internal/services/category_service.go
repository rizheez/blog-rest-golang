package services

import (
	"blog-rest/internal/models"
	"blog-rest/internal/repository"
)

type CategoryService interface {
	GetAllCategories() ([]models.Category, error)
	GetCategoriesByid(id uint) (*models.Category, error)
	CreateCategory(category *models.Category) error
	UpdateCategory(category *models.Category) error
	DeleteCategory(category *models.Category) error
}

func NewCategoryService(c repository.CategoryRepository) CategoryService {
	return &categoryService{categoryRepo: c}
}

type categoryService struct {
	categoryRepo repository.CategoryRepository
}

func (s *categoryService) GetAllCategories() ([]models.Category, error) {
	return s.categoryRepo.GetAllCategories()
}

func (s *categoryService) GetCategoriesByid(id uint) (*models.Category, error) {
	return s.categoryRepo.GetCategoriesByid(id)
}

func (s *categoryService) CreateCategory(category *models.Category) error {
	return s.categoryRepo.CreateCategory(category)
}

func (s *categoryService) UpdateCategory(category *models.Category) error {
	return s.categoryRepo.UpdateCategory(category)
}

func (s *categoryService) DeleteCategory(category *models.Category) error {
	return s.categoryRepo.DeleteCategory(category)
}
