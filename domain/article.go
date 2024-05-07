package domain

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Title       string
	Description string
	Content     string
}
