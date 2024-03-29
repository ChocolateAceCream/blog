package initializer

import (
	"context"
	"errors"
	"fmt"

	"github.com/ChocolateAceCream/blog/global"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"
)

const InitCasbinOrder = InitOrderInternal + 1

type casbinInitializer struct{}

func init() {
	Register(InitCasbinOrder, &casbinInitializer{})
}

func (ci *casbinInitializer) Name() string {
	return "casbin"
}

func (ci *casbinInitializer) Initialize(ctx context.Context) (next context.Context, err error) {
	guestRules := []gormadapter.CasbinRule{
		{Ptype: "p", V1: "/api/v1/user/list", V2: "GET"},
		{Ptype: "p", V1: "/api/v1/user/active", V2: "POST"},
		{Ptype: "p", V1: "/api/v1/user/resetPassword", V2: "PUT"},
		{Ptype: "p", V1: "/api/v1/user/edit", V2: "PUT"},
		{Ptype: "p", V1: "/api/v1/user/delete", V2: "DELETE"},
		{Ptype: "p", V1: "/api/v1/menu/currentUserMenu", V2: "GET"},
		{Ptype: "p", V1: "/api/v1/article/preview", V2: "GET"},
		{Ptype: "p", V1: "/api/v1/article/list", V2: "GET"},
	}
	adminRules := []gormadapter.CasbinRule{
		{Ptype: "p", V1: "/api/v1/role/add", V2: "POST"},
		{Ptype: "p", V1: "/api/v1/role/delete", V2: "DELETE"},
		{Ptype: "p", V1: "/api/v1/role/edit", V2: "PUT"},
		{Ptype: "p", V1: "/api/v1/role/list", V2: "GET"},
		{Ptype: "p", V1: "/api/v1/article/add", V2: "POST"},
		{Ptype: "p", V1: "/api/v1/article/edit", V2: "PUT"},
		{Ptype: "p", V1: "/api/v1/article/delete", V2: "DELETE"},
		{Ptype: "p", V1: "/api/v1/article/search", V2: "GET"},
		{Ptype: "p", V1: "/api/v1/oss/upload", V2: "POST"},
		{Ptype: "p", V1: "/api/v1/comment/add", V2: "POST"},
		{Ptype: "p", V1: "/api/v1/comment/like", V2: "POST"},
		{Ptype: "p", V1: "/api/v1/comment/delete", V2: "DELETE"},
		{Ptype: "p", V1: "/api/v1/comment/list", V2: "GET"},
		{Ptype: "p", V1: "/api/v1/reply/add", V2: "POST"},
		{Ptype: "p", V1: "/api/v1/reply/like", V2: "POST"},
		{Ptype: "p", V1: "/api/v1/reply/delete", V2: "DELETE"},
		{Ptype: "p", V1: "/api/v1/reply/list", V2: "GET"},

		{Ptype: "p", V1: "/api/v1/notification/ws", V2: "GET"},
		{Ptype: "p", V1: "/api/v1/notification/read", V2: "PATCH"},
		{Ptype: "p", V1: "/api/v1/notification/list", V2: "GET"},
		{Ptype: "p", V1: "/api/v1/notification/unreadCount", V2: "GET"},
		{Ptype: "p", V1: "/api/v1/notification/delete", V2: "DELETE"},
	}
	adminRules = append(adminRules, guestRules...)
	superadminRules := []gormadapter.CasbinRule{
		{Ptype: "p", V1: "/api/v1/menu/list", V2: "GET"},
		{Ptype: "p", V1: "/api/v1/menu/delete", V2: "DELETE"},
		{Ptype: "p", V1: "/api/v1/menu/add", V2: "POST"},
		{Ptype: "p", V1: "/api/v1/menu/edit", V2: "PUT"},
		{Ptype: "p", V1: "/api/v1/menu/getRoleMenuTree", V2: "POST"},
		{Ptype: "p", V1: "/api/v1/menu/assignRoleMenus", V2: "POST"},

		{Ptype: "p", V1: "/api/v1/endpoint/list", V2: "GET"},
		{Ptype: "p", V1: "/api/v1/endpoint/add", V2: "POST"},
		{Ptype: "p", V1: "/api/v1/endpoint/delete", V2: "DELETE"},
		{Ptype: "p", V1: "/api/v1/endpoint/edit", V2: "PUT"},

		{Ptype: "p", V1: "/api/v1/casbin/list", V2: "GET"},
		{Ptype: "p", V1: "/api/v1/casbin/update", V2: "POST"},
	}
	superadminRules = append(superadminRules, adminRules...)
	m := map[string][]gormadapter.CasbinRule{
		"1": superadminRules,
		"2": adminRules,
		"3": guestRules,
	}

	db := global.DB
	next = ctx
	for roleId, rules := range m {
		for i := range rules {
			rules[i].V0 = roleId
		}
		if err := db.Create(&rules).Error; err != nil {
			return next, fmt.Errorf("fail to init casbin data, err: %w", err)
		}
		next = context.WithValue(next, ci.Name()+roleId, rules)
	}
	return next, nil
}

func (ci *casbinInitializer) InitDataVerify(ctx context.Context) bool {
	record := gormadapter.CasbinRule{}
	err := global.DB.Where(gormadapter.CasbinRule{Ptype: "p", V0: "1", V1: "/api/v1/menu/add", V2: "GET"}).First(&record).Error
	return !errors.Is(err, gorm.ErrRecordNotFound)
}
