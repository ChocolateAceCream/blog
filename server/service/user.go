package service

import (
	"errors"

	"github.com/ChocolateAceCream/blog/global"
	"github.com/ChocolateAceCream/blog/model/dbTable"
	"github.com/ChocolateAceCream/blog/utils"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type UserService struct{}

func (userService *UserService) GetUserInfoList() (list interface{}, total int64, err error) {
	db := global.DB.Model(&dbTable.User{})
	var userList []dbTable.User
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	// err = db.Limit(limit).Offset(offset).Preload("Authorities").Preload("Authority").Find(&userList).Error
	err = db.Preload("UserRoles").Find(&userList).Error
	return userList, total, err
}

func (userService *UserService) RegisterUser(u dbTable.User) (registeredUser dbTable.User, err error) {
	var user dbTable.User
	if !errors.Is(global.DB.Where("username = ?", u.Username).First(&user).Error, gorm.ErrRecordNotFound) {
		// check if username already taken
		return registeredUser, errors.New("username already taken, please try again")
	} else if !errors.Is(global.DB.Where("email = ?", u.Email).First(&user).Error, gorm.ErrRecordNotFound) {
		// check if email already taken
		return registeredUser, errors.New("email already taken, please try again")
	}

	u.Password = utils.BcryptHash(u.Password)
	u.UUID = uuid.NewV4()
	err = global.DB.Create(&u).Error
	return u, err
}

func (userService *UserService) EditUser(u dbTable.User) error {
	return global.DB.Model(&dbTable.User{}).Where("UUID = ? ", u.UUID).Updates(&u).Error
}

func (userService *UserService) DeleteUser(uuid uuid.UUID) error {
	return global.DB.Model(&dbTable.User{}).Where("uuid = ?", uuid).Delete(&dbTable.User{}).Error
}
