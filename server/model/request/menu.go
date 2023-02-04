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
			Type:      1, //menu item
			Meta: dbTable.Meta{
				Icon:  "home",
				Title: "主页",
			},
		},
		{
			MODEL:     global.MODEL{ID: 2},
			Pid:       1,
			Name:      "homeChild1",
			Component: "@/views/home/child1",
			Path:      "/home/child1",
			Type:      1, //menu item
			Meta: dbTable.Meta{
				Icon:  "home",
				Title: "主页child1",
			},
		},
		{
			MODEL:     global.MODEL{ID: 2},
			Pid:       1,
			Name:      "homeChild2",
			Component: "@/views/home/child2",
			Path:      "/home/child2",
			Type:      1, //menu item
			Meta: dbTable.Meta{
				Icon:  "home",
				Title: "主页child2",
			},
		},
	}
}
