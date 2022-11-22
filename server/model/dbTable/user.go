package dbTable

import (
	"github.com/ChocolateAceCream/blog/global"
	uuid "github.com/satori/go.uuid"
)

type User struct {
	global.MODEL
	UUID     uuid.UUID `json:"uuid" gorm:"comment:UUID" mapstructure:"UUID"`
	Username string    `json:"username" gorm:"comment:username;unique" mapstructure:"Username"`
	Password string    `json:"password" gorm:"comment:password" mapstructure:"Password"`
	Email    string    `json:"email" gorm:"comment:email;unique" mapstructure:"Email"`
	// UserRole  Role      `json:"role" gorm:"foreignKey:RoleId;references:RoleId;comment:user's role"`
	Active    int    `json:"active" gorm:"default:1;comment:if user is active 1 active 2 inactive" mapstructure:"Active"`
	UserRoles []Role `json:"userRoles" gorm:"many2many:userRole" mapstructure:"UserRoles"`
}
