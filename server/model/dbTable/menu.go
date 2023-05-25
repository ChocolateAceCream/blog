package dbTable

import "time"

type Menu struct {
	ID        uint      `json:"id" gorm:"primarykey"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Pid       uint      `json:"pid" gorm:"comment:parent id"`
	Path      string    `json:"path" gorm:"comment:route path, base menu start with /"`
	Name      string    `json:"name" gorm:"comment:route name"`
	Component string    `json:"component" gorm:"comment:web component path, e.g. @/views/auth/register"`
	Display   int       `json:"display" gorm:"comment:hide on menu: 1-show on menu, 2-hide"`
	Meta      `json:"meta" gorm:"embedded"`
	Roles     []Role `json:"roles" gorm:"many2many:roleMenu;constraint:OnDelete:CASCADE;"`
	// ChildMenu []Menu `gorm:"foreignkey:Pid;constraint:OnDelete:CASCADE;"`
}

type Meta struct {
	Icon      string `json:"icon" gorm:"comment:menu icon name, no suffix"`
	Title     string `json:"title" gorm:"comment:breadcrumb title"`
	KeepAlive int    `json:"keepAlive" gorm:"default:1;comment:1-keep alive, 2-not keep alive"`
}
