package service

import (
	"github.com/ChocolateAceCream/blog/global"
	"github.com/ChocolateAceCream/blog/model/dbTable"
)

type ArticleService struct{}

func (articleService *ArticleService) GetArticleInfo(articleId int) (a dbTable.Article, err error) {
	err = global.DB.Preload("Author").Where("id = ?", articleId).First(&a).Error
	return
}

func (*ArticleService) AddArticle(a dbTable.Article) (dbTable.Article, error) {
	err := global.DB.Create(&a).Error
	return a, err
}

func (*ArticleService) EditArticle(a dbTable.Article) error {
	return global.DB.Model(&dbTable.Article{}).Where("ID = ? ", a.ID).Updates(&a).Error
}

func (*ArticleService) HasPermission(authorId uint, articleId uint) bool {
	r := dbTable.Article{}
	err := global.DB.First(&r, articleId).Error
	if err != nil || r.AuthorID != authorId {
		return false
	} else {
		return true
	}
}
