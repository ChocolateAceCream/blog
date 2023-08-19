package dbTable

import "time"

type ReplyLiker struct {
	ID        uint      `json:"id" gorm:"primarykey"`
	CreatedAt time.Time `json:"createdAt"`
	UserID    uint      `gorm:"uniqueIndex: idx_ReplyLiker"`
	ReplyID   uint      `gorm:"uniqueIndex: idx_ReplyLiker"`
}
