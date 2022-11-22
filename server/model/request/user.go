package request

import uuid "github.com/satori/go.uuid"

// User register structure
type RegisterUser struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	// HeaderImg    string `json:"headerImg" gorm:"default:'https://qmplusimg.henrongyi.top/gva_header.jpg'"`
	Email string `json:"email" binding:"required"`
	// Captcha string `json:"captcha" binding:"required"`  // use email verification instead
	RoleId  uint   `json:"role" gorm:"default:888"`
	Active  int    `json:"active"`
	RoleIds []uint `json:"roles"`
}

type ActiveUser struct {
	Code int `json:"code" binding:"required"`
}

type EditUser struct {
	UUID     uuid.UUID `json:"uuid" gorm:"primarykey" binding:"required"` // uuid
	Username string    `json:"username"`
	RoleId   uint      `json:"role" gorm:"default:888"`
	Email    string    `json:"email"  `
	Active   int       `json:"active"`
	// Authorities []model.SysAuthority `json:"-" gorm:"many2many:sys_user_authority;"`
}
type DeleteUser struct {
	UUID uuid.UUID `json:"uuid" gorm:"primarykey" binding:"required"` // uuid
}
