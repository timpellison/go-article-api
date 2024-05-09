package domain

import (
	"gorm.io/gorm"
	"time"
)

type Article struct {
	gorm.Model
	Title       string
	Description string
	Content     string
	PublishDate time.Time
}
