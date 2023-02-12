package dbTable

import "github.com/ChocolateAceCream/blog/global"

type Endpoint struct {
	global.MODEL
	Path        string `json:"path" gorm:"comment:server api endpoint" binding:"required_if=Type 2"`
	Description string `json:"description" gorm:"comment: api endpoint description"`
	Method      string `json:"method" gorm:"default:POST;comment:endpoint call method" binding:"required_if=Type 2"`
	PID         uint   `json:"pid" gorm:"comment:parent group id" binding:"required"`
	Name        string `json:"name" gorm:"comment: displayed name" binding:"required"`
	Type        int    `json:"type" gorm:"default:1;comment:1-group, 2-endpoint" binding:"oneof=1 2"`
}
