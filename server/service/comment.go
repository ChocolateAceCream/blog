package service

import (
	"errors"
	"fmt"

	"github.com/ChocolateAceCream/blog/global"
	"github.com/ChocolateAceCream/blog/model/dbTable"
	"github.com/ChocolateAceCream/blog/model/request"
	"github.com/ChocolateAceCream/blog/model/response"
	"github.com/ChocolateAceCream/blog/utils"
)

type CommentService struct{}

func (*CommentService) AddComment(c dbTable.Comment, articleId uint) error {
	var article dbTable.Article
	fmt.Println("-----------c.ArticleID-----", c.ArticleID)
	if err := global.DB.First(&article, c.ArticleID).Error; err != nil {
		fmt.Println("-----------err-----", err)
		return errors.New("article not found")
	}
	return global.DB.Create(&c).Error
}

func (es *CommentService) GetCommentList(query request.CommentCursorListParam) (commentBaseInfo []response.CommentBaseInfo, total int64, err error) {
	db := global.DB.Model(&dbTable.Comment{}).Where("article_id = ? ", query.ArticleID)
	commentList := []dbTable.Comment{}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	fmt.Println("----total-----", total)
	db = db.Limit(query.PageSize)
	queryStr := "id > ?"
	if query.Desc {
		db = db.Order("id desc")
		queryStr = "id < ?"
	}
	if query.CursorId > 0 {
		db = db.Where(queryStr, query.CursorId)
	}
	err = db.Preload("Author").Preload("Replies").Find(&commentList).Error //TODO: preload with conditions: only preload certain number of replies
	commentBaseInfo = utils.MapSlice(commentList, response.CommentBaseInfoFormatter)
	return
}
func (commentService *CommentService) DeleteComment(authorId uint, ids []int) (err error) {
	//TODO: test delete associated role-comment relations
	comments := []dbTable.Comment{}
	if err = global.DB.Where("author_id = ? AND id in ?", authorId, ids).Delete(&comments).Error; err != nil {
		return err
	}
	return nil
}
