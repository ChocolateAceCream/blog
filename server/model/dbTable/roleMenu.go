package dbTable

type RoleMenu struct {
	RoleId uint `gorm:"column:role_id"`
	MenuId uint `gorm:"column:menu_id"`
}
