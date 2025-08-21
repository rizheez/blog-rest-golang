package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title      string   `json:"title"`
	Body       string   `json:"body"`
	UserID     uint     `json:"user_id"`
	User       User     `json:"user"`
	CategoryID uint     `json:"category_id"`
	Category   Category `json:"category"`
}
