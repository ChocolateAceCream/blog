package model

import (
	"github.com/ChocolateAceCream/blog/global"
	uuid "github.com/satori/go.uuid"
)

type User struct {
	global.MODEL
	UUID     uuid.UUID `json:"uuid" gorm:"comment:UUID"`
	Username string    `json:"username" gorm:"comment:username"`
	Password string    `json:"password" gorm:"comment:password"`
	Email    string    `json:"email" gorm:"comment:email"`
	// UserRole  Role      `json:"role" gorm:"foreignKey:RoleId;references:RoleId;comment:user's role"`
	UserRoles []Role `json:"userRoles" gorm:"many2many:userRole"`
}
