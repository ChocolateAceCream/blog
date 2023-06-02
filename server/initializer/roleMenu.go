package initializer

import (
	"context"
	"fmt"

	"github.com/ChocolateAceCream/blog/global"
	"github.com/ChocolateAceCream/blog/model/dbTable"
)

const InitRoleMenuOrder = InitRoleOrder + InitMenuOrder

type initRoleMenu struct{}

func init() {
	Register(InitRoleMenuOrder, &initRoleMenu{})
}

func (rim *initRoleMenu) Name() string {
	return "roleMenu"
}

func (rim *initRoleMenu) Initialize(ctx context.Context) (next context.Context, err error) {
	ri := roleInitializer{}
	rm := menuInitializer{}
	db := global.DB
	next = ctx
	for _, name := range []string{"superadmin", "admin", "guest"} {
		menus, ok := next.Value(rm.Name() + name).([]dbTable.Menu)
		if !ok {
			return next, fmt.Errorf("fail to find %s menu in role-menu initializer", name)
		}
		role, ok := next.Value(ri.Name() + name).(dbTable.Role)
		if !ok {
			return next, fmt.Errorf("fail to find %s role in role-menu initializer", name)
		}
		if err = db.Model(&role).Association("Menus").Append(menus); err != nil {
			return next, err
		}
	}
	return next, nil
}

func (rim *initRoleMenu) InitDataVerify(ctx context.Context) bool {
	var record dbTable.Role
	r := global.DB.Where("name = ? ", "superadmin").Preload("Menus").First(&record)
	if r != nil && r.Error == nil {
		return len(record.Menus) > 0
	}
	return false
}
