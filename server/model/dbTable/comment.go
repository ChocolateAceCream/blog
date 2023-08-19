package dbTable

import "github.com/ChocolateAceCream/blog/global"

type Comment struct {
	global.MODEL
	Author         User           `json:"author" gorm:"foreignKey:AuthorID"`
	AuthorID       uint           `json:"authorId" gorm:"comment:foreignKey" binding:"required"`
	CommentContent string         `json:"commentContent" gorm:"type:text;comment:article comment content" binding:"required"`
	LikesCount     uint           `json:"likesCount" gorm:"default:0;comment:comment likes count" binding:"required"`
	Replies        []Reply        `json:"replies" gorm:"foreignKey:CommentID"`
	RepliesCount   uint           `json:"repliesCount" gorm:"default:0;comment:replies count" `
	ArticleID      uint           `json:"articleID" gorm:"comment:foreignKey" binding:"required"`
	CommentLikers  []CommentLiker `gorm:"foreignKey:CommentID;"`
	IsLiked        bool           `json:"isLiked"`
}
