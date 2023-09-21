package dbTable

import "github.com/ChocolateAceCream/blog/global"

type Notification struct {
	global.MODEL
	Status      uint   `json:"status" gorm:"comment:notification status: 1-read, 2-unread"`
	Type        uint   `json:"type" gorm:"comment:type of notification: 1-comment likes, 2-article reply, 3-comment reply"`
	Initiator   User   `json:"initiator" gorm:"foreignKey:InitiatorID"`
	InitiatorID uint   `json:"initiatorID" gorm:"comment:foreignKey"`
	Recipient   User   `json:"recipient" gorm:"foreignKey:RecipientID"`
	RecipientID uint   `json:"recipientID" gorm:"comment:foreignKey"`
	Content     string `json:"content" gorm:"comment:Notification content"`
}
