package service

import (
	"errors"
	"strconv"

	"github.com/ChocolateAceCream/blog/global"
	"github.com/ChocolateAceCream/blog/model/dbTable"
	"gorm.io/gorm"
)

type RoleService struct{}

type AuthorityService struct{}

var AuthorityServiceApp = new(AuthorityService)

func (roleService *RoleService) AddRole(payload dbTable.Role) (dbTable.Role, error) {
	var role dbTable.Role
	if !errors.Is(global.DB.Where("name = ?", payload.Name).First(&role).Error, gorm.ErrRecordNotFound) {
		return role, errors.New("role already exists")
	}
	err := global.DB.Create(&payload).Error
	return payload, err
}

func (roleService *RoleService) GetUserRole(u dbTable.User) (dbTable.Role, error) {
	var roles []dbTable.Role
	err := global.DB.Model(&u).Association("Roles").Find(&roles)
	return roles[0], err
}

func (roleService *RoleService) EditRole(r dbTable.Role) error {
	return global.DB.Model(&dbTable.Role{}).Where("ID = ? ", r.ID).Updates(&r).Error
}

func (roleService *RoleService) GetRoleList() (roles []dbTable.Role, err error) {
	//TODO: filter out unused field by define GetRoleMenus response struct
	err = global.DB.Find(&roles).Error
	return roles, err
}

func (roleService *RoleService) DeleteRole(id []int) (err error) {
	//TODO: test delete associated role-menu relations
	roles := []dbTable.Role{}
	uintId := []uint{}
	for _, v := range id {
		roles = append(roles, dbTable.Role{ID: uint(v)})
		uintId = append(uintId, uint(v))
	}

	tx := global.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err = tx.Error; err != nil {
		return err
	}
	if err = tx.Select("Menus", "Users").Delete(&roles).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err = tx.Commit().Error; err != nil {
		return err
	}
	for _, v := range id {
		roleId := strconv.Itoa(v)
		var casbinService CasbinService
		casbinService.ClearCasbin(0, roleId) //role id is index 0 in casbin
	}
	return err
}
