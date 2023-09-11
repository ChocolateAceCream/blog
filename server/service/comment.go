package service

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/ChocolateAceCream/blog/global"
	"github.com/ChocolateAceCream/blog/library"
	"github.com/ChocolateAceCream/blog/model/dbTable"
	"github.com/ChocolateAceCream/blog/model/request"
	"github.com/ChocolateAceCream/blog/model/response"
	"github.com/ChocolateAceCream/blog/utils"
	"gorm.io/gorm"
)

type CommentService struct{}

func (*CommentService) LikeComment(p request.LikeCommentPayload) error {
	library.PublishMqttMsg(strconv.FormatUint(uint64(p.UserID), 10), strconv.FormatUint(uint64(p.UserID), 10))
	if *p.Like {
		r := global.DB.Create(&dbTable.CommentLiker{UserID: p.UserID, CommentID: p.CommentID})
		fmt.Println(r.RowsAffected == 1)
		if r.RowsAffected == 1 {
			return global.DB.Model(&dbTable.Comment{}).Where("id = ?", p.CommentID).UpdateColumn("likes_count", gorm.Expr("likes_count + ?", 1)).Error
		}
		return r.Error
	} else {
		r := global.DB.Where("user_id = ? and comment_id = ?", p.UserID, p.CommentID).Delete(&dbTable.CommentLiker{})
		fmt.Println(r.RowsAffected == 1)
		if r.RowsAffected == 1 {
			return global.DB.Model(&dbTable.Comment{}).Where("id = ?", p.CommentID).UpdateColumn("likes_count", gorm.Expr("likes_count - ?", 1)).Error
		}
		return r.Error
	}
}

func (*CommentService) AddComment(c dbTable.Comment, articleId uint) error {
	var article dbTable.Article
	fmt.Println("-----------c.ArticleID-----", c.ArticleID)
	if err := global.DB.First(&article, c.ArticleID).Error; err != nil {
		fmt.Println("-----------err-----", err)
		return errors.New("article not found")
	}
	return global.DB.Create(&c).Error
}

func (es *CommentService) GetCommentList(query request.CommentCursorListParam, currentUser dbTable.User) (commentBaseInfo []response.CommentBaseInfo, total int64, err error) {
	db := global.DB.Model(&dbTable.Comment{}).Where("article_id = ? ", query.ArticleID)
	commentList := []dbTable.Comment{}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	fmt.Println("----total-----", total)
	db = db.Limit(query.PageSize)
	queryStr := "comments.id > ?"
	if query.Desc {
		db = db.Order("id desc")
		queryStr = "comments.id < ?"
	}
	if query.CursorId > 0 {
		db = db.Where(queryStr, query.CursorId)
	}
	subQuery := global.DB.Table("comment_likers").
		Select("comment_id, COUNT(*) > 0 AS is_liked").
		Where("user_id = ?", currentUser.ID).
		Group("comment_id")
	db.
		Select("comments.*, COALESCE(subquery.is_liked, false) AS is_liked").
		Joins("LEFT JOIN (?) AS subquery ON comments.id = subquery.comment_id", subQuery).
		Joins("Author").
		// Joins("Replies").
		Scan(&commentList)

	// err = db.Preload("Author").Preload("Replies").Find(&commentList).Error //TODO: preload with conditions: only preload certain number of replies
	commentBaseInfo = utils.MapSlice(commentList, response.CommentBaseInfoFormatter)
	return
}
func (commentService *CommentService) DeleteComment(authorId uint, id int) (err error) {
	//TODO: test delete associated role-comment relations
	var comment dbTable.Comment
	comment.ID = uint(id)
	if err = global.DB.Select("Replies").Where("author_id = ? ", authorId).Delete(&comment).Error; err != nil {
		return err
	}
	return nil
}
