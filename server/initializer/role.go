// TODO: add role users?
package initializer

import (
	"context"
	"errors"
	"fmt"

	"github.com/ChocolateAceCream/blog/global"
	"github.com/ChocolateAceCream/blog/model/dbTable"
	"gorm.io/gorm"
)

const InitRoleOrder = InitCasbinOrder + 1

type roleInitializer struct{}

func init() {
	Register(InitRoleOrder, &roleInitializer{})
}

func (ri *roleInitializer) Name() string {
	return "role"
}

func (ri *roleInitializer) Initialize(ctx context.Context) (next context.Context, err error) {
	db := global.DB
	entities := []dbTable.Role{
		{
			ID:   1,
			Name: "superadmin",
			Pid:  0,
		},
		{
			ID:   2,
			Name: "admin",
			Pid:  1,
		},
		{
			ID:   3,
			Name: "guest",
			Pid:  2,
		},
	}
	if err = db.Create(&entities).Error; err != nil {
		return ctx, fmt.Errorf("fail to init role data, err: %w", err)
	}
	next = ctx
	for _, e := range entities {
		next = context.WithValue(next, ri.Name()+e.Name, e)
	}
	return next, nil

}

func (ri *roleInitializer) InitDataVerify(ctx context.Context) bool {
	var record dbTable.Role
	err := global.DB.Where("name = ? ", "superadmin").First(&record).Error
	return !errors.Is(err, gorm.ErrRecordNotFound)
}
