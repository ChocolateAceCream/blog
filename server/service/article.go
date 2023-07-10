package service

import (
	"github.com/ChocolateAceCream/blog/global"
	"github.com/ChocolateAceCream/blog/model/dbTable"
	"github.com/ChocolateAceCream/blog/model/request"
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

func (es *ArticleService) GetArticleList(query request.ArticleSearchParma) (articleList []dbTable.Article, total int64, err error) {
	db := global.DB.Model(&dbTable.Article{})
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	db = db.Limit(query.PageSize)
	queryStr := "id > ?"
	if query.Desc {
		db = db.Order("id desc")
		queryStr = "id < ?"
	}
	if query.CursorId > 0 {
		db = db.Where(queryStr, query.CursorId)
	}
	err = db.Preload("Author").Find(&articleList).Error
	return
}
