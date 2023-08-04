package dbTable

import "github.com/ChocolateAceCream/blog/global"

type Reply struct {
	global.MODEL
	Author          User   `json:"author" gorm:"foreignKey:AuthorID"`
	AuthorID        uint   `json:"authorId" gorm:"comment:foreignKey" binding:"required"`
	ReplyContent    string `json:"replyContent" gorm:"type:text;comment: reply content of article comment" binding:"required"`
	ParentReplyID   uint   `json:"parentReplyID" gorm:"default:0;comment: parent reply id if exist, otherwise is 0" `
	ParentCommentID uint   `json:"parentCommentID" gorm:"comment: parent comment id" binding:"required"`
	LikesCount      uint   `json:"likesCount" gorm:"default:0;comment:comment likes count" binding:"required"`
}
