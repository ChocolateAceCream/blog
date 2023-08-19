package service

import (
	"errors"

	"github.com/ChocolateAceCream/blog/global"
	"github.com/ChocolateAceCream/blog/model/dbTable"
	"github.com/ChocolateAceCream/blog/model/request"
	"github.com/ChocolateAceCream/blog/model/response"
	"github.com/ChocolateAceCream/blog/utils"
	"gorm.io/gorm"
)

type ReplyService struct{}

func (*ReplyService) AddReply(c dbTable.Reply) error {
	var comment dbTable.Comment
	if err := global.DB.First(&comment, c.CommentID).Error; err != nil {
		return errors.New("comment not found")
	}
	if c.ParentReplyID != nil {
		var reply dbTable.Reply
		if err := global.DB.First(&reply, *c.ParentReplyID).Error; err != nil {
			return errors.New("parent reply not found")
		}
	}
	if err := global.DB.Create(&c).Error; err != nil {
		return err
	}

	return global.DB.Model(&dbTable.Comment{}).Where("id = ?", c.CommentID).UpdateColumn("replies_count", gorm.Expr("replies_count + ?", 1)).Error
}

func (replyService *ReplyService) DeleteReply(authorId uint, id int) (err error) {
	var reply dbTable.Reply
	if err := global.DB.First(&reply, id).Error; err != nil {
		return errors.New("reply not found")
	}
	if reply.AuthorID != authorId {
		return errors.New("reply author not matching current user")
	}
	commentID := reply.CommentID
	if err = global.DB.Delete(&reply).Error; err != nil {
		return err
	}
	//TODO: test commentID after delete?
	return global.DB.Model(&dbTable.Comment{}).Where("id = ?", commentID).UpdateColumn("replies_count", gorm.Expr("replies_count - ?", 1)).Error
}

func (es *ReplyService) GetReplyList(query request.ReplyCursorListParam, currentUser dbTable.User) (replyBaseInfo []response.ReplyBaseInfo, total int64, err error) {
	db := global.DB.Model(&dbTable.Reply{}).Where("comment_id = ? ", query.CommentID)
	replyList := []dbTable.Reply{}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	db = db.Limit(query.PageSize)
	queryStr := "replies.id > ?"
	if query.Desc {
		db = db.Order("id desc")
		queryStr = "replies.id < ?"
	}
	if query.CursorId > 0 {
		db = db.Where(queryStr, query.CursorId)
	}
	subQuery := global.DB.Table("reply_likers").
		Select("reply_id, COUNT(*) > 0 AS is_liked").
		Where("user_id = ?", currentUser.ID).
		Group("reply_id")
	db.
		Select("replies.*, COALESCE(subquery.is_liked, false) AS is_liked").
		Joins("LEFT JOIN (?) AS subquery ON replies.id = subquery.reply_id", subQuery).
		Joins("Author").
		// Joins("Replies").
		Scan(&replyList).
		// Preload("ParentReply.Author").Find(&replyList)

		Preload("ParentReply.Author", func(db *gorm.DB) *gorm.DB {
			return db.Select("username,id") // always remember to select the foreign key, which in this case, is id
			//TODO: if avatar needed, remember to edit here to include avatar
		}).
		Find(&replyList)

	// err = db.Preload("Author").Preload("Replies").Find(&replyList).Error //TODO: preload with conditions: only preload certain number of replies
	replyBaseInfo = utils.MapSlice(replyList, response.ReplyBaseInfoFormatter)
	return
}

// func (*ReplyService) LikeReply(p request.LikeReplyPayload) error {
// 	if *p.Like {
// 		r := global.DB.Create(&dbTable.ReplyLiker{UserID: p.UserID, ReplyID: p.ReplyID})
// 		fmt.Println(r.RowsAffected == 1)
// 		if r.RowsAffected == 1 {
// 			return global.DB.Model(&dbTable.Reply{}).Where("id = ?", p.ReplyID).UpdateColumn("likes_count", gorm.Expr("likes_count + ?", 1)).Error
// 		}
// 		return r.Error
// 	} else {
// 		r := global.DB.Where("user_id = ? and reply_id = ?", p.UserID, p.ReplyID).Delete(&dbTable.ReplyLiker{})
// 		fmt.Println(r.RowsAffected == 1)
// 		if r.RowsAffected == 1 {
// 			return global.DB.Model(&dbTable.Reply{}).Where("id = ?", p.ReplyID).UpdateColumn("likes_count", gorm.Expr("likes_count - ?", 1)).Error
// 		}
// 		return r.Error
// 	}
// }
