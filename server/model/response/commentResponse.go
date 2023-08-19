package response

import (
	"time"

	"github.com/ChocolateAceCream/blog/model/dbTable"
)

type CommentBaseInfo struct {
	UpdatedAt    time.Time `json:"updatedAt"`
	Content      string    `json:"content"`
	AuthorID     uint      `json:"authorId"`
	Author       string    `json:"author"`
	ID           uint      `json:"id"`
	LikesCount   uint      `json:"likesCount"`
	IsLiked      bool      `json:"isLiked"`
	RepliesCount uint      `json:"repliesCount" `
}

func CommentBaseInfoFormatter(a dbTable.Comment) CommentBaseInfo {
	return CommentBaseInfo{
		UpdatedAt:    a.UpdatedAt,
		Content:      a.CommentContent,
		AuthorID:     a.AuthorID,
		Author:       a.Author.Username,
		ID:           a.ID,
		LikesCount:   a.LikesCount,
		IsLiked:      a.IsLiked,
		RepliesCount: a.RepliesCount,
	}
}
