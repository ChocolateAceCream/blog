package request

type NotificationCursorListQuery struct {
	Params NotificationCursorListParam `json:"params" form:"params" binding:"required"`
}

type NotificationCursorListParam struct {
	CursorListParam
	UnreadOnly bool `json:"unreadOnly" form:"unreadOnly" default:"false"`
}
