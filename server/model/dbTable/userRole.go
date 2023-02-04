package dbTable

type UserRole struct {
	UserId uint `gorm:"column:user_id"`
	RoleId uint `gorm:"column:role_id"`
}
