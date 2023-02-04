package initializer

import (
	"context"
	"errors"
	"fmt"

	"github.com/ChocolateAceCream/blog/global"
	"github.com/ChocolateAceCream/blog/model/dbTable"
	"github.com/ChocolateAceCream/blog/utils"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

const InitUserOrder = InitRoleOrder + 1

type userInitilizer struct{}

func init() {
	Register(InitUserOrder, &userInitilizer{})
}

func (ui *userInitilizer) Name() string {
	return "user"
}

func (ui *userInitilizer) Initialize(ctx context.Context) (next context.Context, err error) {
	config := global.CONFIG.Init
	adminPassword := utils.BcryptHash(config.AdminEmail)
	guestPassword := utils.BcryptHash(config.GuestPassword)
	adminEmail, guestEmail := config.AdminEmail, config.GuestEmail
	db := global.DB
	entities := []dbTable.User{
		{
			UUID:     uuid.NewV4(),
			Username: "superadmin",
			Password: adminPassword,
			Email:    "super@super.com",
			RoleId:   1,
			Active:   1,
		},
		{
			UUID:     uuid.NewV4(),
			Username: "admin",
			Password: adminPassword,
			Email:    adminEmail,
			RoleId:   2,
			Active:   1,
		},
		{
			UUID:     uuid.NewV4(),
			Username: "guest",
			Password: guestPassword,
			Email:    guestEmail,
			RoleId:   3,
			Active:   1,
		},
	}
	if err = db.Create(&entities).Error; err != nil {
		return ctx, fmt.Errorf("fail to init user data, err: %w", err)
	}
	next = context.WithValue(ctx, ui.Name(), entities)
	ri := roleInitilizer{}
	for _, user := range entities {
		role, ok := ctx.Value(ri.Name() + fmt.Sprint(user.RoleId)).(dbTable.Role)
		if !ok {
			return next, errors.New("fail to associate user with role")
		}
		if err = db.Model(&user).Association("Role").Append(&role); err != nil {
			return next, err
		}
	}
	return next, err
}

func (ui *userInitilizer) DataInitialized(ctx context.Context) bool {
	var record dbTable.User
	err := global.DB.Where("username = ? ", "superadmin").Preload("Role").First(&record).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false
	}
	return record.Role.ID == 1
}
