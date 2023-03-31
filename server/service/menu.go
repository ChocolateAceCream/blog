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
	//TODO: filter out unused field by define GetRoleMenus response struct
	err = global.DB.Model(&dbTable.Role{ID: uint(roleId)}).Association("Menus").Find(&menus)
	return menus, err
}

func (menuService MenuService) GetMenuList() (menus []dbTable.Menu, err error) {
	//TODO: filter out unused field by define GetRoleMenus response struct
	err = global.DB.Find(&menus).Error
	return menus, err
}

func (menuService MenuService) DeleteMenu(id []int) (err error) {
	//TODO: test delete associated role-menu relations
	menus := []dbTable.Menu{}
	for _, v := range id {
		menus = append(menus, dbTable.Menu{ID: uint(v)})
	}
	return global.DB.Select("Roles").Delete(&menus).Error
	// return global.DB.Where("id = ?", id).Delete(&dbTable.Menu{}).Error
}

func (menuService MenuService) EditMenu(m dbTable.Menu) error {
	return global.DB.Model(&dbTable.Menu{}).Where("ID = ? ", m.ID).Updates(&m).Error
}
