package persistence

import (
	"fmt"
	"gorm.io/gorm"
	"goservice/config"
	"goservice/domain"
)

type IArticleRepository interface {
	Add(article *domain.Article) *domain.Article
	Delete(article *domain.Article)
	GetOne(id uint) *domain.Article
	GetMany() []domain.Article
}

type ArticleRepository struct {
	database *gorm.DB
	config   *config.Configuration
}

func NewArticleRepository(configuration *config.Configuration, database *gorm.DB) (result *ArticleRepository) {
	result = &ArticleRepository{database: database, config: configuration}
}

func (r ArticleRepository) Add(article *domain.Article) *domain.Article {
	id := r.database.Create(article)
	r.database.First(&article, id)
	return article
}

func (r ArticleRepository) Delete(article *domain.Article) {
	fmt.Printf("Article to delete is %v", article)
	r.database.Delete(&article)
}

func (r ArticleRepository) GetOne(id uint) *domain.Article {
	var article domain.Article
	fmt.Printf("Article id to get is %v", id)

	r.database.First(&article, id)
	if article.ID == 0 {
		return nil
	}
	return &article
}

func (r ArticleRepository) GetMany() *[]domain.Article {
	var articles []domain.Article
	r.database.Find(&articles)
	return &articles
}
