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
const GROUP = 1
const ENDPOINT = 2

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
		{
			MODEL:       global.MODEL{ID: 1},
			PID:         0,
			Description: "user related endpoints",
			Name:        "User Group",
			Type:        GROUP,
		},
		{
			MODEL:       global.MODEL{ID: 11},
			PID:         1,
			Description: "get user list",
			Name:        "Get User List",
			Method:      "GET",
			Path:        "/api/v1/user/userList",
			Type:        ENDPOINT,
		},
		{
			MODEL:       global.MODEL{ID: 12},
			PID:         1,
			Description: "Active User",
			Name:        "Active User",
			Method:      "POST",
			Path:        "/api/v1/user/active",
			Type:        ENDPOINT,
		},
		{
			MODEL:       global.MODEL{ID: 13},
			PID:         1,
			Description: "Reset user password",
			Name:        "Reset Password",
			Method:      "PUT",
			Path:        "/api/v1/user/resetPassword",
			Type:        ENDPOINT,
		},
		{
			MODEL:       global.MODEL{ID: 14},
			PID:         1,
			Description: "Edit user info",
			Name:        "Edit User",
			Method:      "PUT",
			Path:        "/api/v1/user/edit",
			Type:        ENDPOINT,
		},
		{
			MODEL:       global.MODEL{ID: 15},
			PID:         1,
			Description: "Delete user ",
			Name:        "Delete User",
			Method:      "DELETE",
			Path:        "/api/v1/user/delete",
			Type:        ENDPOINT,
		},

		{
			MODEL:       global.MODEL{ID: 2},
			PID:         0,
			Description: "Role related endpoints",
			Name:        "Role Group",
			Type:        GROUP,
		},
		{
			MODEL:       global.MODEL{ID: 21},
			PID:         2,
			Description: "Create Role",
			Name:        "Create Role",
			Method:      "POST",
			Path:        "/api/v1/role/create",
			Type:        ENDPOINT,
		},
		{
			MODEL:       global.MODEL{ID: 3},
			PID:         0,
			Description: "Menu related endpoints",
			Name:        "Menu Group",
			Type:        GROUP,
		},
		{
			MODEL:       global.MODEL{ID: 31},
			PID:         3,
			Description: "Create Meu",
			Name:        "Create Menu",
			Method:      "POST",
			Path:        "/api/v1/menu/create",
			Type:        ENDPOINT,
		},
		{
			MODEL:       global.MODEL{ID: 32},
			PID:         3,
			Description: "Get Current Menu",
			Name:        "Get Current Menu",
			Method:      "GET",
			Path:        "/api/v1/menu/currentUserMenu",
			Type:        ENDPOINT,
		},
		{
			MODEL:       global.MODEL{ID: 4},
			PID:         0,
			Description: "Casbin related endpoints",
			Name:        "Casbin Group",
			Type:        GROUP,
		},
		{
			MODEL:       global.MODEL{ID: 5},
			PID:         0,
			Description: "Endpoint related endpoints",
			Name:        "Endpoint Group",
			Type:        GROUP,
		},
	}
	if err = db.Create(&entities).Error; err != nil {
		return ctx, fmt.Errorf("fail to init endpoint data, err: %w", err)
	}
	next = context.WithValue(ctx, ei.Name(), entities)
	return next, nil
}

func (ei *endpointInitilizer) InitDataVerify(ctx context.Context) bool {
	record := dbTable.Endpoint{}
	err := global.DB.Where(dbTable.Endpoint{Name: "User Group"}).First(&record).Error
	return !errors.Is(err, gorm.ErrRecordNotFound)
}
