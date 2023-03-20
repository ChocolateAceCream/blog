package initializer

import (
	"context"
	"errors"
	"fmt"

	"github.com/ChocolateAceCream/blog/global"
	"github.com/ChocolateAceCream/blog/model/dbTable"
	"gorm.io/gorm"
)

const InitMenuOrder = InitRoleOrder + 1

type menuInitilizer struct{}

func init() {
	Register(InitMenuOrder, &menuInitilizer{})
}

func (ri *menuInitilizer) Name() string {
	return "menu"
}

func (ri *menuInitilizer) Initialize(ctx context.Context) (next context.Context, err error) {
	db := global.DB
	guestMenus := []dbTable.Menu{
		{
			MODEL:     global.MODEL{ID: 1},
			Pid:       0,
			Name:      "home",
			Component: "views/home/index.vue",
			Path:      "/home",
			Meta: dbTable.Meta{
				Icon:  "home",
				Title: "主页",
			},
		},
	}
	adminMenus := []dbTable.Menu{
		{
			MODEL:     global.MODEL{ID: 2},
			Pid:       0,
			Name:      "admin",
			Component: "views/admin/index.vue",
			Path:      "/admin",
			Meta: dbTable.Meta{
				Icon:  "admin",
				Title: "管理员",
			},
		},
		{
			MODEL:     global.MODEL{ID: 3},
			Pid:       2,
			Name:      "role",
			Component: "views/admin/role/index.vue",
			Path:      "role",
			Meta: dbTable.Meta{
				Icon:  "role",
				Title: "角色管理",
			},
		},
	}
	adminMenus = append(adminMenus, guestMenus...)
	superadminMenus := []dbTable.Menu{
		{
			MODEL:     global.MODEL{ID: 4},
			Pid:       2,
			Name:      "menu",
			Component: "views/admin/menu/index.vue",
			Path:      "menu",
			Meta: dbTable.Meta{
				Icon:  "menu",
				Title: "菜单管理",
			},
		},
		{
			MODEL:     global.MODEL{ID: 5},
			Pid:       2,
			Name:      "endpoint",
			Component: "views/admin/endpoint/index.vue",
			Path:      "endpoint",
			Meta: dbTable.Meta{
				Icon:  "endpoint",
				Title: "api管理",
			},
		},
	}
	superadminMenus = append(superadminMenus, adminMenus...)
	if err = db.Create(&superadminMenus).Error; err != nil {
		return ctx, fmt.Errorf("fail to init menu data, err: %w", err)
	}
	next = ctx
	next = context.WithValue(next, ri.Name()+"superadmin", superadminMenus)
	next = context.WithValue(next, ri.Name()+"admin", adminMenus)
	next = context.WithValue(next, ri.Name()+"guest", guestMenus)
	return next, nil
}

func (ri *menuInitilizer) InitDataVerify(ctx context.Context) bool {
	err := global.DB.Where("name = ?", "home").First(&dbTable.Menu{}).Error
	return !errors.Is(err, gorm.ErrRecordNotFound)
}
