package service

import (
	"errors"

	"github.com/ChocolateAceCream/blog/global"
	"github.com/ChocolateAceCream/blog/model/dbTable"
	"github.com/ChocolateAceCream/blog/model/response"
	"gorm.io/gorm"
)

type ArticleService struct{}

func (articleService *ArticleService) GetArticleInfo(id int) (article response.ArticleInfo, err error) {
	var a dbTable.Article
	if err := global.DB.Preload("Author").Where("id = ?", id).First(&a).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return article, errors.New("article not found ")
	}

	fileBytes, err := global.FS.ReadFile(global.CONFIG.Local.StaticFilePath + a.Path)
	article = response.ArticleInfo{
		File:    fileBytes,
		Article: a,
	}
	return
}
