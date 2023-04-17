package initializer

import (
	"context"
	"errors"
	"fmt"

	"github.com/ChocolateAceCream/blog/global"
	"github.com/ChocolateAceCream/blog/model/dbTable"
	"gorm.io/gorm"
)

const InitEndpointOrder = InitOrderInternal + 1

type endpointInitilizer struct{}

func init() {
	Register(InitEndpointOrder, &endpointInitilizer{})
}

func (ei *endpointInitilizer) Name() string {
	return "endpoint"
}
func (ei *endpointInitilizer) Initialize(ctx context.Context) (next context.Context, err error) {
	db := global.DB
	entities := []dbTable.Endpoint{
		{Method: "GET", Path: "/api/v1/user/list", Group: "User", Description: "Get user list", Name: "Get User List"},
		{Method: "POST", Path: "/api/v1/user/active", Group: "User", Description: "Active User", Name: "Active User"},
		{Method: "PUT", Path: "/api/v1/user/resetPassword", Group: "User", Description: "Reset user password", Name: "Reset Password"},
		{Method: "PUT", Path: "/api/v1/user/edit", Group: "User", Description: "Edit user info", Name: "Edit User"},
		{Method: "DELETE", Path: "/api/v1/user/delete", Group: "User", Description: "Delete user ", Name: "Delete User"},

		{Method: "POST", Path: "/api/v1/role/add", Group: "Role", Description: "Create Role", Name: "Create Role"},
		{Method: "DELETE", Path: "/api/v1/role/delete", Group: "Role", Description: "Delete Role", Name: "Delete Role"},
		{Method: "PUT", Path: "/api/v1/role/edit", Group: "Role", Description: "Edit Role", Name: "Edit Role"},
		{Method: "GET", Path: "/api/v1/role/list", Group: "Role", Description: "Get Role List", Name: "Get Role List"},

		{Method: "POST", Path: "/api/v1/menu/create", Group: "Menu", Description: "Create Menu", Name: "Create Menu"},
		{Method: "GET", Path: "/api/v1/menu/currentUserMenu", Group: "Menu", Description: "Get current user's menu ", Name: "Get Current User Menu"},
		{Method: "GET", Path: "/api/v1/menu/list", Group: "Menu", Description: "Get all Menu list", Name: "List Menu"},
		{Method: "DELETE", Path: "/api/v1/menu/delete", Group: "Menu", Description: "Delete Menu", Name: "Delete Menu"},
		{Method: "POST", Path: "/api/v1/menu/add", Group: "Menu", Description: "Add Menu", Name: "Add Menu"},
		{Method: "PUT", Path: "/api/v1/menu/edit", Group: "Menu", Description: "Edit Menu", Name: "Edit Menu"},
		{Method: "POST", Path: "/api/v1/menu/getRoleMenuTree", Group: "Menu", Description: "Get Role Menu Tree", Name: "Role Menu Tree"},
		{Method: "POST", Path: "/api/v1/menu/assignRoleMenus", Group: "Menu", Description: "Assign Role Menus", Name: "Assign Role Menus"},

		{Method: "GET", Path: "/api/v1/endpoint/list", Group: "Endpoint", Description: "Get Endpoint list", Name: "List Endpoint"},
		{Method: "POST", Path: "/api/v1/endpoint/add", Group: "Endpoint", Description: "Add Endpoint", Name: "Add Endpoint"},
		{Method: "PUT", Path: "/api/v1/endpoint/edit", Group: "Endpoint", Description: "Edit Endpoint", Name: "Edit Endpoint"},
		{Method: "DELETE", Path: "/api/v1/endpoint/delete", Group: "Endpoint", Description: "Delete Endpoint", Name: "Delete Endpoint"},
	}
	if err = db.Create(&entities).Error; err != nil {
		return ctx, fmt.Errorf("fail to init endpoint data, err: %w", err)
	}
	next = context.WithValue(ctx, ei.Name(), entities)
	return next, nil
}

func (ei *endpointInitilizer) InitDataVerify(ctx context.Context) bool {
	record := dbTable.Endpoint{}
	err := global.DB.Where(dbTable.Endpoint{Name: "Get User List"}).First(&record).Error
	return !errors.Is(err, gorm.ErrRecordNotFound)
}
