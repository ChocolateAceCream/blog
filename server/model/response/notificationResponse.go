package response

import (
	"time"

	"github.com/ChocolateAceCream/blog/model/dbTable"
)

type NotificationBaseInfo struct {
	UpdatedAt   time.Time `json:"updatedAt"`
	Content     string    `json:"content"`
	InitiatorID uint      `json:"initiatorId"`
	Initiator   string    `json:"initiator"`
	ID          uint      `json:"id"`
	Type        string    `json:"type"`
	Status      string    `json:"status"`
}

func NotificationBaseInfoFormatter(a dbTable.Notification) NotificationBaseInfo {

	typeMapper := map[uint]string{
		1: "likeComment",
		2: "articleReply",
		3: "commentReply",
	}

	statusMapper := map[uint]string{
		1: "read",
		2: "unread",
	}

	return NotificationBaseInfo{
		UpdatedAt:   a.UpdatedAt,
		Content:     a.Content,
		Initiator:   a.Initiator.Username,
		InitiatorID: a.InitiatorID,
		ID:          a.ID,
		Status:      statusMapper[a.Status],
		Type:        typeMapper[a.Type],
	}
}
