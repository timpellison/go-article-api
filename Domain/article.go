package Domain

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Id          int32
	Title       string
	Description string
	Content     string
}
