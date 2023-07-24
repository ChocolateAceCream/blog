package service

import (
	"github.com/ChocolateAceCream/blog/global"
	"github.com/ChocolateAceCream/blog/model/dbTable"
	"github.com/ChocolateAceCream/blog/model/request"
	"github.com/ChocolateAceCream/blog/model/response"
	"github.com/ChocolateAceCream/blog/utils"
	"gorm.io/gorm"
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

func (*ArticleService) ViewedTimesPlusOne(articleId int) {
	global.DB.Model(&dbTable.Article{}).Where("ID = ? ", articleId).UpdateColumn("viewed_times", gorm.Expr("viewed_times + ?", 1))
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

func (es *ArticleService) GetArticleList(query request.ArticleCursorListParma) (articleBaseInfo []response.ArticleBaseInfo, total int64, err error) {
	db := global.DB.Model(&dbTable.Article{})
	articleList := []dbTable.Article{}
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
	articleBaseInfo = utils.MapSlice(articleList, response.ArticleBaseInfoFormatter)
	return
}

func (es *ArticleService) GetArticleSearchList(authorId uint, query request.ArticleSearchParma) (articleBaseInfo []response.ArticleBaseInfo, total int64, err error) {
	// fmt.Println("----query----", query)
	articleList := []dbTable.Article{}
	sql := `
	SELECT SQL_CALC_FOUND_ROWS * FROM articles a
	WHERE 1=1
	AND
	deleted_at IS NULL
	AND
	author_id = ?
	AND
	published = ?
	AND
	CONCAT_WS('', a.content, a.title) LIKE CONCAT('%', ?, '%')
	ORDER BY created_at DESC
	LIMIT ? OFFSET ?
	`
	limit := query.PageSize
	offset := query.PageSize * (query.PageNumber - 1)
	db := global.DB.Debug().Raw(sql, authorId, query.Published, query.Keywords, limit, offset).Scan(&articleList)
	// fmt.Println("--------articleList---", articleList)
	err = db.Debug().Raw("select found_rows() as count").Scan(&total).Error
	if err != nil {
		return
	}
	articleBaseInfo = utils.MapSlice(articleList, response.ArticleBaseInfoFormatter)
	return
}

func (articleService *ArticleService) DeleteArticle(authorId uint, ids []int) (err error) {
	//TODO: test delete associated role-article relations
	articles := []dbTable.Article{}
	if err = global.DB.Where("author_id = ? AND id in ?", authorId, ids).Delete(&articles).Error; err != nil {
		return err
	}
	return nil
}
