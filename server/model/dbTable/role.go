package dbTable

import "time"

type Role struct {
	CreatedAt time.Time  // 创建时间
	UpdatedAt time.Time  // 更新时间
	DeletedAt *time.Time `sql:"index"`
	ID        uint       `json:"roleId" gorm:"primarykey;comment:role ID;size:90;"`
	Name      string     `json:"name" gorm:"not null;unique;comment:role name" binding:"required"`
	Users     []User     `json:"users"`
	ParentId  uint       `json:"parentId" gorm:"comment:parent role id" binding:"required"`
	Children  []Role     `json:"children" gorm:"-"`
	Menus     []Menu     `json:"-" gorm:"many2many:roleMenu"`
}
