package initializer

import (
	"context"
	"fmt"

	"github.com/ChocolateAceCream/blog/global"
	"github.com/ChocolateAceCream/blog/model/dbTable"
)

const InitUserRoleOrder = InitRoleOrder + InitUserOrder

type initUserRole struct{}

func init() {
	Register(InitUserRoleOrder, &initUserRole{})
}

func (rim *initUserRole) Name() string {
	return "userRole"
}

func (rim *initUserRole) Initialize(ctx context.Context) (next context.Context, err error) {
	ri := roleInitilizer{}
	ru := userInitilizer{}
	db := global.DB
	next = ctx
	for _, name := range []string{"superadmin", "admin", "guest"} {
		user, ok := next.Value(ru.Name() + name).(dbTable.User)
		if !ok {
			return next, fmt.Errorf("fail to find %s user in user-role initilizer", name)
		}
		role, ok := next.Value(ri.Name() + name).(dbTable.Role)
		if !ok {
			return next, fmt.Errorf("fail to find %s role in user-role initilizer", name)
		}
		if err = db.Model(&user).Association("Roles").Append([]dbTable.Role{role}); err != nil {
			return next, err
		}
	}
	return next, nil
}

func (rim *initUserRole) InitDataVerify(ctx context.Context) bool {
	var record dbTable.Role
	r := global.DB.Where("name = ? ", "superadmin").Preload("Users").First(&record)
	if r != nil && r.Error == nil {
		return len(record.Users) > 0
	}
	return false
}
