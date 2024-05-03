package Persistence

import "goservice/Domain"

type IArticleRepository interface {
	func Add(article *Domain.Article) *Domain.Article
	func Delete(article *Domain.Article)
	func GetOne(id int32) *Domain.Article
}

func Add(article *Domain.Article) (*Domain.Article) {

}