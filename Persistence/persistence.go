package Persistence

import (
	"fmt"
	"gorm.io/gorm"
	"goservice/Config"
	"goservice/Domain"
)

type IArticleRepository interface {
	Add(article *Domain.Article) *Domain.Article
	Delete(article *Domain.Article)
	GetOne(id uint) *Domain.Article
	GetMany() []Domain.Article
}

type ArticleRepository struct {
	database *gorm.DB
	config   *Config.Configuration
}

func NewArticleRepository(configuration *Config.Configuration, database *gorm.DB) *ArticleRepository {
	result := &ArticleRepository{database: database, config: configuration}
	return result
}

func (r ArticleRepository) Add(article *Domain.Article) *Domain.Article {
	id := r.database.Create(article)
	r.database.First(&article, id)
	return article
}

func (r ArticleRepository) Delete(article *Domain.Article) {
	fmt.Printf("Article to delete is %v", article)
	r.database.Delete(&article)
}

func (r ArticleRepository) GetOne(id uint) *Domain.Article {
	var article Domain.Article
	fmt.Printf("Article id to get is %v", id)

	r.database.First(&article, id)
	if article.ID == 0 {
		return nil
	}
	return &article
}

func (r ArticleRepository) GetMany() *[]Domain.Article {
	var articles []Domain.Article
	r.database.Find(&articles)
	return &articles
}
