package request

import (
	"github.com/ChocolateAceCream/blog/model/dbTable"
)

type AssignRoleMenus struct {
	FindById
	Menus []dbTable.Menu `json:"menus" binding:"required"`
}

func DefaultMenus() []dbTable.Menu {
	return []dbTable.Menu{
		{
			ID:        1,
			Pid:       0,
			Name:      "home",
			Component: "@/views/home",
			Path:      "/home",
			Display:   1,
			Meta: dbTable.Meta{
				Icon:  "home",
				Title: "主页",
			},
		},
	}
}
