package repository

import (
	"blog-rest/internal/models"

	"gorm.io/gorm"
)

type PostRepository interface {
	GetAllPosts() ([]models.Post, error)
	GetPostById(id int) (models.Post, error)
	CreatePost(post *models.Post) error
	UpdatePost(post *models.Post) error
	DeletePost(post *models.Post) error
}

type postRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return &postRepository{db: db}
}

func (r *postRepository) GetAllPosts() ([]models.Post, error) {
	var posts []models.Post
	result := r.db.Preload("User").Preload("Category").Find(&posts)
	err := result.Error
	return posts, err
}

func (r *postRepository) GetPostById(id int) (models.Post, error) {
	post := models.Post{}
	result := r.db.Preload("User").Preload("Category").First(&post, id)
	err := result.Error
	return post, err
}

func (r *postRepository) CreatePost(post *models.Post) error {
	return r.db.Create(post).Error
}

func (r *postRepository) UpdatePost(post *models.Post) error {
	return r.db.Model(&models.Post{}).Where("id = ?", post.ID).Updates(post).Error
}

func (r *postRepository) DeletePost(post *models.Post) error {
	return r.db.Delete(post).Error
}
