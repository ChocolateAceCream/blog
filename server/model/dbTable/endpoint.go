package dbTable

import "github.com/ChocolateAceCream/blog/global"

type Endpoint struct {
	global.MODEL
	Path        string `json:"path" gorm:"comment:server api endpoint" binding:"required"`
	Description string `json:"description" gorm:"comment: api endpoint description"`
	Method      string `json:"method" gorm:"default:POST;comment:endpoint call method" binding:"required"`
	Name        string `json:"name" gorm:"comment: displayed name" binding:"required"`
	GroupName   string `json:"groupName" gorm:"comment:endpoint group, to display endpoints by group"`
}
