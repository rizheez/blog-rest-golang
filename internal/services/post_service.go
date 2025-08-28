package services

import (
	"blog-rest/internal/models"
	"blog-rest/internal/repository"
)

type PostService interface {
	GetPosts() ([]models.Post, error)
	GetPostById(id int) (models.Post, error)
	CreatePost(post *models.Post) error
	UpdatePost(post *models.Post) error
	DeletePost(post *models.Post) error
}

type postService struct {
	postRepo repository.PostRepository
}

func NewPostService(postRepo repository.PostRepository) PostService {
	return &postService{postRepo: postRepo}
}

func (s *postService) GetPosts() ([]models.Post, error) {
	return s.postRepo.GetAllPosts()
}

func (s *postService) GetPostById(id int) (models.Post, error) {
	return s.postRepo.GetPostById(id)
}

func (s *postService) CreatePost(post *models.Post) error {
	return s.postRepo.CreatePost(post)
}

func (s *postService) UpdatePost(post *models.Post) error {
	return s.postRepo.UpdatePost(post)
}

func (s *postService) DeletePost(post *models.Post) error {
	return s.postRepo.DeletePost(post)
}
