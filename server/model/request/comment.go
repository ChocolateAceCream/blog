package request

type CommentCursorListQuery struct {
	Params CommentCursorListParam `json:"params" form:"params" binding:"required"`
}

type CommentCursorListParam struct {
	CursorListParam
	ArticleID uint `json:"articleId" form:"articleId" binding:"required"`
}

type AddCommentPayload struct {
	CommentContent string `json:"commentContent" form:"commentContent" binding:"required"`
	ArticleID      uint   `json:"articleId" form:"articleId" binding:"required"`
}

type LikeCommentPayload struct {
	CommentID uint  `json:"commentId" form:"commentId" binding:"required"`
	Like      *bool `json:"like" form:"like" binding:"required"`
	UserID    uint
}
