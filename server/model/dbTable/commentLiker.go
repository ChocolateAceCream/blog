package dbTable

import (
	"time"
)

type CommentLiker struct {
	ID        uint      `json:"id" gorm:"primarykey"`
	CreatedAt time.Time `json:"createdAt"`
	UserID    uint      `gorm:"uniqueIndex: idx_CommentLiker"`
	CommentID uint      `gorm:"uniqueIndex: idx_CommentLiker"`
}
