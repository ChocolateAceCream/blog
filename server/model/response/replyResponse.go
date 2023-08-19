package response

import (
	"time"

	"github.com/ChocolateAceCream/blog/model/dbTable"
)

type ReplyBaseInfo struct {
	UpdatedAt     time.Time      `json:"updatedAt"`
	Content       string         `json:"content"`
	AuthorID      uint           `json:"authorId"`
	Author        string         `json:"author"`
	ID            uint           `json:"id"`
	LikesCount    uint           `json:"likesCount"`
	IsLiked       bool           `json:"isLiked"`
	ParentReplyID *uint          `json:"parentReplyId,omitempty"`
	ParentReply   *dbTable.Reply `json:"parentReply,omitempty"`
}

func ReplyBaseInfoFormatter(a dbTable.Reply) ReplyBaseInfo {
	r := ReplyBaseInfo{
		UpdatedAt:  a.UpdatedAt,
		Content:    a.ReplyContent,
		AuthorID:   a.AuthorID,
		Author:     a.Author.Username,
		ID:         a.ID,
		LikesCount: a.LikesCount,
		IsLiked:    a.IsLiked,
	}
	if *a.ParentReplyID != 0 {
		r.ParentReply = a.ParentReply
		r.ParentReplyID = a.ParentReplyID
	}
	return r
}
