package service

import (
	"github.com/ChocolateAceCream/blog/db/model"
	"github.com/ChocolateAceCream/blog/global"
)

type UserService struct{}

func (userService *UserService) GetUserInfoList() (list interface{}, total int64, err error) {
	db := global.DB.Model(&model.User{})
	var userList []model.User
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	// err = db.Limit(limit).Offset(offset).Preload("Authorities").Preload("Authority").Find(&userList).Error
	err = db.Preload("UserRoles").Find(&userList).Error
	return userList, total, err
}
