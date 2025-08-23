package dto

type CreatePostRequest struct {
	Title      string `json:"title" validate:"required"`
	Body       string `json:"body" validate:"required"`
	UserID     uint   `json:"user_id" validate:"required"`
	CategoryID uint   `json:"category_id" validate:"required"`
}

type UpdatePostRequest struct {
	Title      string `json:"title"`
	Body       string `json:"body"`
	CategoryID uint   `json:"category_id"`
}

type PostResponse struct {
	ID        uint             `json:"id"`
	Title     string           `json:"title"`
	Body      string           `json:"body"`
	User      UserResponse     `json:"user"`
	Category  CategoryResponse `json:"category"`
	CreatedAt string           `json:"created_at"`
	UpdatedAt string           `json:"updated_at"`
}
