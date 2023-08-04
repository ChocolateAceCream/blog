package response

import (
	"time"

	"github.com/ChocolateAceCream/blog/model/dbTable"
)

type CommentBaseInfo struct {
	UpdatedAt  time.Time `json:"updatedAt"`
	Content    string    `json:"content"`
	AuthorID   uint      `json:"authorId"`
	Author     string    `json:"author"`
	ID         uint      `json:"id"`
	LikesCount uint      `json:"likesCount"`
}

func CommentBaseInfoFormatter(a dbTable.Comment) CommentBaseInfo {
	return CommentBaseInfo{
		UpdatedAt:  a.MODEL.UpdatedAt,
		Content:    a.CommentContent,
		AuthorID:   a.AuthorID,
		Author:     a.Author.Username,
		ID:         a.ID,
		LikesCount: a.LikesCount,
	}
}
