package response

import "github.com/ChocolateAceCream/blog/model/dbTable"

type RoleMenuTree struct {
	MenuList  []dbTable.Menu `json:"menuList"`
	RoleMenus []dbTable.Menu `json:"roleMenus"`
}
