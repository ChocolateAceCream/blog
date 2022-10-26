package dbTable

import "time"

type Role struct {
	CreatedAt time.Time  // 创建时间
	UpdatedAt time.Time  // 更新时间
	DeletedAt *time.Time `sql:"index"`
	RoleId    uint       `json:"roleId" gorm:"not null;unique;primary_key;comment:role ID;size:90"`
	RoleName  string     `json:"roleName" gorm:"not null;unique;comment:role name"`
	Users     []User     `json:"-" gorm:"many2many:userRole"`
}
