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

type endpointInitializer struct{}

func init() {
	Register(InitEndpointOrder, &endpointInitializer{})
}

func (ei *endpointInitializer) Name() string {
	return "endpoint"
}
func (ei *endpointInitializer) Initialize(ctx context.Context) (next context.Context, err error) {
	db := global.DB
	entities := []dbTable.Endpoint{
		{Method: "GET", Path: "/api/v1/user/list", GroupName: "User", Description: "Get user list", Name: "Get User List"},
		{Method: "POST", Path: "/api/v1/user/active", GroupName: "User", Description: "Active User", Name: "Active User"},
		{Method: "PUT", Path: "/api/v1/user/resetPassword", GroupName: "User", Description: "Reset user password", Name: "Reset Password"},
		{Method: "PUT", Path: "/api/v1/user/edit", GroupName: "User", Description: "Edit user info", Name: "Edit User"},
		{Method: "DELETE", Path: "/api/v1/user/delete", GroupName: "User", Description: "Delete user ", Name: "Delete User"},

		{Method: "POST", Path: "/api/v1/role/add", GroupName: "Role", Description: "Create Role", Name: "Create Role"},
		{Method: "DELETE", Path: "/api/v1/role/delete", GroupName: "Role", Description: "Delete Role", Name: "Delete Role"},
		{Method: "PUT", Path: "/api/v1/role/edit", GroupName: "Role", Description: "Edit Role", Name: "Edit Role"},
		{Method: "GET", Path: "/api/v1/role/list", GroupName: "Role", Description: "Get Role List", Name: "Get Role List"},

		{Method: "GET", Path: "/api/v1/menu/currentUserMenu", GroupName: "Menu", Description: "Get current user's menu ", Name: "Get Current User Menu"},
		{Method: "GET", Path: "/api/v1/menu/list", GroupName: "Menu", Description: "Get all Menu list", Name: "List Menu"},
		{Method: "DELETE", Path: "/api/v1/menu/delete", GroupName: "Menu", Description: "Delete Menu", Name: "Delete Menu"},
		{Method: "POST", Path: "/api/v1/menu/add", GroupName: "Menu", Description: "Add Menu", Name: "Add Menu"},
		{Method: "PUT", Path: "/api/v1/menu/edit", GroupName: "Menu", Description: "Edit Menu", Name: "Edit Menu"},
		{Method: "POST", Path: "/api/v1/menu/getRoleMenuTree", GroupName: "Menu", Description: "Get Role Menu Tree", Name: "Role Menu Tree"},
		{Method: "POST", Path: "/api/v1/menu/assignRoleMenus", GroupName: "Menu", Description: "Assign Role Menus", Name: "Assign Role Menus"},

		{Method: "GET", Path: "/api/v1/endpoint/list", GroupName: "Endpoint", Description: "Get Endpoint list", Name: "List Endpoint"},
		{Method: "POST", Path: "/api/v1/endpoint/add", GroupName: "Endpoint", Description: "Add Endpoint", Name: "Add Endpoint"},
		{Method: "PUT", Path: "/api/v1/endpoint/edit", GroupName: "Endpoint", Description: "Edit Endpoint", Name: "Edit Endpoint"},
		{Method: "DELETE", Path: "/api/v1/endpoint/delete", GroupName: "Endpoint", Description: "Delete Endpoint", Name: "Delete Endpoint"},

		{Method: "GET", Path: "/api/v1/article/preview", GroupName: "Article", Description: "Get article content", Name: "Preview  Article"},
		{Method: "GET", Path: "/api/v1/article/list", GroupName: "Article", Description: "Get article list", Name: "List  Article"},
		{Method: "GET", Path: "/api/v1/article/search", GroupName: "Article", Description: "Get article search list", Name: "search  Article"},
		{Method: "POST", Path: "/api/v1/article/add", GroupName: "Article", Description: "Add Article", Name: "Add Article"},
		{Method: "PUT", Path: "/api/v1/article/edit", GroupName: "Article", Description: "Edit Article", Name: "Edit Article"},
		{Method: "DELETE", Path: "/api/v1/article/delete", GroupName: "Article", Description: "Delete Endpoint", Name: "Delete Endpoint"},

		{Method: "POST", Path: "/api/v1/casbin/update", GroupName: "Casbin", Description: "Update Role's Casbin", Name: "Update Casbin"},
		{Method: "GET", Path: "/api/v1/casbin/list", GroupName: "Casbin", Description: "Get casbin list", Name: "list Casbin"},
	}
	if err = db.Create(&entities).Error; err != nil {
		return ctx, fmt.Errorf("fail to init endpoint data, err: %w", err)
	}
	next = context.WithValue(ctx, ei.Name(), entities)
	return next, nil
}

func (ei *endpointInitializer) InitDataVerify(ctx context.Context) bool {
	record := dbTable.Endpoint{}
	err := global.DB.Where(dbTable.Endpoint{Name: "Get User List"}).First(&record).Error
	return !errors.Is(err, gorm.ErrRecordNotFound)
}
