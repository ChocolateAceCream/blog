package request

import (
	"github.com/ChocolateAceCream/blog/global"
	"github.com/ChocolateAceCream/blog/model/dbTable"
)

func DefaultMenus() []dbTable.Menu {
	return []dbTable.Menu{
		{
			MODEL:     global.MODEL{ID: 1},
			Pid:       0,
			Name:      "home",
			Component: "@/views/home",
			Path:      "/home",
			Meta: dbTable.Meta{
				Icon:  "home",
				Title: "主页",
			},
		},
	}
}
