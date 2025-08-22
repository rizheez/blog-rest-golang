package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name  string `json:"name"`
	Posts []Post `json:"-"`
}
