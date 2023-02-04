package dbTable

import (
	"github.com/ChocolateAceCream/blog/global"
)

type Menu struct {
	global.MODEL
	Pid       uint   `json:"pid" gorm:"comment:parent id"`
	Path      string `json:"path" gorm:"comment:route path, base menu start with /"`
	Name      string `json:"name" gorm:"comment:route name"`
	Component string `json:"component" gorm:"comment:web component path, e.g. @/views/auth/register"`
	Meta      `json:"meta" gorm:"embedded"`
	Type      int    `json:"type" gorm:"default:1;comment:1-menu item, 2-button"`
	Roles     []Role `json:"roles" gorm:"many2many:roleMenu"`
}

type Meta struct {
	Icon      string `json:"icon" gorm:"comment:menu icon name, no suffix"`
	Title     string `json:"title" gorm:"comment:breadcrumb title"`
	KeepAlive int    `json:"keepAlive" gorm:"default:1;comment:1-keep alive, 2-not keep alive"`
}
