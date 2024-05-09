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

func NewArticleRepository(configuration *config.Configuration, database *gorm.DB) (*ArticleRepository, error) {
	result := &ArticleRepository{database: database, config: configuration}
	return result, nil
}

func (r ArticleRepository) Add(article *domain.Article) (*domain.Article, error) {
	var id = r.database.Create(article)
	if id.Error != nil {
		return nil, id.Error
	}
	r.database.First(&article, id)
	return article, nil
}

func (r ArticleRepository) Delete(article *domain.Article) error {
	fmt.Printf("Article to delete is %v", article)
	tx := r.database.Delete(&article)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (r ArticleRepository) GetOne(id uint) (*domain.Article, error) {
	var article domain.Article
	fmt.Printf("Article id to get is %v", id)

	tx := r.database.First(&article, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &article, nil
}

func (r ArticleRepository) GetMany() (*[]domain.Article, error) {
	var articles []domain.Article
	tx := r.database.Find(&articles)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &articles, nil
}
