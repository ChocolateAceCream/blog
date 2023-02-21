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

type casbinInitilizer struct{}

func init() {
	Register(InitCasbinOrder, &casbinInitilizer{})
}

func (ci *casbinInitilizer) Name() string {
	return "casbin"
}

func (ci *casbinInitilizer) Initialize(ctx context.Context) (next context.Context, err error) {
	guestRules := []gormadapter.CasbinRule{
		{Ptype: "p", V1: "/api/v1/user/userList", V2: "GET"},
		{Ptype: "p", V1: "/api/v1/user/active", V2: "POST"},
		{Ptype: "p", V1: "/api/v1/user/resetPassword", V2: "PUT"},
		{Ptype: "p", V1: "/api/v1/user/edit", V2: "PUT"},
		{Ptype: "p", V1: "/api/v1/user/delete", V2: "DELETE"},
	}
	adminRules := []gormadapter.CasbinRule{
		{Ptype: "p", V1: "/api/v1/role/create", V2: "POST"},
	}
	adminRules = append(adminRules, guestRules...)
	superadminRules := []gormadapter.CasbinRule{
		{Ptype: "p", V1: "/api/v1/menu/create", V2: "POST"},
		{Ptype: "p", V1: "/api/v1/menu/currentUserMenu", V2: "GET"},
		{Ptype: "p", V1: "/api/v1/endpoint/all", V2: "GET"},
		{Ptype: "p", V1: "/api/v1/endpoint/list", V2: "POST"},
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

func (ci *casbinInitilizer) InitDataVerify(ctx context.Context) bool {
	record := gormadapter.CasbinRule{}
	err := global.DB.Where(gormadapter.CasbinRule{Ptype: "p", V0: "1", V1: "/api/v1/menu/create", V2: "GET"}).First(&record).Error
	return !errors.Is(err, gorm.ErrRecordNotFound)
}