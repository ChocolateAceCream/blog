package request

type AddReplyPayload struct {
	ReplyContent  string `json:"replyContent" form:"replyContent" binding:"required"`
	CommentID     uint   `json:"commentID" form:"commentID" binding:"required"`
	ParentReplyID *uint  `json:"parentReplyId" form:"parentReplyId" binding:"omitempty"`
}

type ReplyCursorListQuery struct {
	Params ReplyCursorListParam `json:"params" form:"params" binding:"required"`
}

type ReplyCursorListParam struct {
	CursorListParam
	CommentID uint `json:"commentId" form:"commentId" binding:"required"`
}

type LikeReplyPayload struct {
	ReplyID uint  `json:"replyId" form:"replyId" binding:"required"`
	Like    *bool `json:"like" form:"like" binding:"required"`
	UserID  uint
}
