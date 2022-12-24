package service

import (
	"errors"

	"github.com/ChocolateAceCream/blog/global"
	"github.com/ChocolateAceCream/blog/model/dbTable"
	"gorm.io/gorm"
)

type RoleService struct{}

type AuthorityService struct{}

var AuthorityServiceApp = new(AuthorityService)

func (roleService *RoleService) CreateRole(payload dbTable.Role) (dbTable.Role, error) {
	var role dbTable.Role
	if !errors.Is(global.DB.Where("name = ?", payload.Name).First(&role).Error, gorm.ErrRecordNotFound) {
		return role, errors.New("role already exists")
	}
	err := global.DB.Create(&payload).Error
	return payload, err
}
