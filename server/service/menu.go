package service

import (
	"errors"

	"github.com/ChocolateAceCream/blog/global"
	"github.com/ChocolateAceCream/blog/model/dbTable"
	"gorm.io/gorm"
)

type MenuService struct{}

func (menuService MenuService) AddMenu(menu dbTable.Menu) error {
	if !errors.Is(global.DB.Where("name = ?", menu.Name).First(&dbTable.Menu{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("menu name already exists, please change")
	}
	return global.DB.Create(&menu).Error
}

func (menuService MenuService) AssignRoleMenus(menus []dbTable.Menu, RoleId uint) error {
	var r dbTable.Role
	global.DB.Preload("Menu").First(&r, "id = ?", RoleId)
	err := global.DB.Model(&r).Association("Menus").Replace(menus)
	return err
}

func (menuService MenuService) GetRoleMenus(roleId int) (menus []dbTable.Menu, err error) {
	var r dbTable.Role
	err = global.DB.Model(&r).Where("id = ?", roleId).Association("Menus").Find(&menus)
	return menus, err
}
