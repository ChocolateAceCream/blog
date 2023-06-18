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

type userInitializer struct{}

func init() {
	Register(InitUserOrder, &userInitializer{})
}

func (ui *userInitializer) Name() string {
	return "user"
}

func (ui *userInitializer) Initialize(ctx context.Context) (next context.Context, err error) {
	config := global.CONFIG.Init
	adminPassword := utils.BcryptHash(config.AdminPassword)
	guestPassword := utils.BcryptHash(config.GuestPassword)
	adminEmail, guestEmail := config.AdminEmail, config.GuestEmail
	db := global.DB
	entities := []dbTable.User{
		{
			UUID:     uuid.NewV4(),
			Username: "superadmin",
			Password: adminPassword,
			Email:    "super@super.com",
			Active:   1,
		},
		{
			UUID:     uuid.NewV4(),
			Username: "admin",
			Password: adminPassword,
			Email:    adminEmail,
			Active:   1,
		},
		{
			UUID:     uuid.NewV4(),
			Username: "guest",
			Password: guestPassword,
			Email:    guestEmail,
			Active:   1,
		},
	}
	if err = db.Create(&entities).Error; err != nil {
		return ctx, fmt.Errorf("fail to init user data, err: %w", err)
	}

	if err = db.Model(&entities[0]).Association("Followers").Append([]dbTable.User{entities[1], entities[2], entities[2]}); err != nil {
		return ctx, err
	}

	/*
		// duplicated record won't be append to associations
		e.g.
		if err = db.Model(&entities[0]).Association("Followers").Append([]dbTable.User{entities[1], entities[2], entities[2]}); err != nil {
			return ctx, err
		}
		//will not insert records since they already exited in db
	*/

	next = ctx
	for _, e := range entities {
		next = context.WithValue(next, ui.Name()+e.Username, e)
	}
	return next, nil
}

func (ui *userInitializer) InitDataVerify(ctx context.Context) bool {
	var record dbTable.User
	err := global.DB.Where("username = ? ", "superadmin").First(&record).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false
	}
	return record.Email == "super@super.com"
}
