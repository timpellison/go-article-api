package persistence

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"goservice/config"
	"goservice/domain"
	"strconv"
	"time"
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

const DatabaseCluster = "DATABASE_CLUSTER"
const DatabaseUserName = "DATABASE_USERNAME"
const DatabasePassword = "DATABASE_PASSWORD"
const DatabaseHost = "DATABASE_PORT"
const DatabaseName = "DATABASE_DATABASENAME"
const DatabasePort = "DATABASE_PORT"

func NewArticleRepository(configuration *config.Configuration) (*ArticleRepository, error) {
	// now for some persistence
	dsn := "host=" + configuration.Database.Cluster +
		" user=" + configuration.Database.UserName +
		" port=" + strconv.FormatInt(int64(configuration.Database.Port), 10) +
		" dbname=" + configuration.Database.DatabaseName +
		" password=" + configuration.Database.Password +
		" sslmode=disable" +
		" TimeZone=UTC"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		for i := 0; i < 10; i++ {
			db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
			if err == nil {
				break
			}
			time.Sleep(1 * time.Second)
		}
		if err != nil {
			panic(err.Error())
		}
	}
	fmt.Printf("CONNECTION SUCCESSFUL!  Connected to database %s on cluster %s", configuration.Database.DatabaseName, configuration.Database.Cluster)
	err = db.AutoMigrate(&domain.Article{})
	if err != nil {
		fmt.Printf("Unable to migrate database, %v", err)
		return nil, err
	}
	result := &ArticleRepository{database: db, config: configuration}
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
