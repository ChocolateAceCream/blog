package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/ChocolateAceCream/blog/global"
	"github.com/ChocolateAceCream/blog/model/dbTable"
	"github.com/ChocolateAceCream/blog/model/request"
	"github.com/ChocolateAceCream/blog/utils"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type UserService struct{}

func (userService *UserService) GetUserInfoList(query request.UserSearchQuery) (userList []dbTable.User, total int64, err error) {
	db := global.DB.Model(&dbTable.User{})

	if query.Active != 0 {
		db = db.Where("active =?", query.Active)
	}
	if query.Username != "" {
		db = db.Where("username LIKE ?", "%"+query.Username+"%")
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	limit := query.PageSize
	offset := query.PageSize * (query.PageNumber - 1)
	db = db.Limit(limit).Offset(offset)
	if query.OrderBy != "" {
		var orderStr string
		// sql injection protection
		orderMap := make(map[string]bool, 3)
		orderMap["username"] = true
		orderMap["email"] = true
		orderMap["id"] = true
		if orderMap[query.OrderBy] {
			if query.Desc {
				orderStr = query.OrderBy + " desc"
			} else {
				orderStr = query.OrderBy
			}
		} else { // didn't matched any order key in `orderMap`
			err = fmt.Errorf("invalid orderby key: %v", query.OrderBy)
			return
		}
		// Offset(-1).Limit(-1) is used to cancel offset and limit effect.
		err = db.Order(orderStr).Find(&userList).Offset(-1).Limit(-1).Error
	} else {
		err = db.Order("id").Find(&userList).Offset(-1).Limit(-1).Error
	}
	return userList, total, err
}

func (userService *UserService) ActiveUser(u dbTable.User, code string) error {
	key := global.CONFIG.Email.Prefix + u.UUID.String()
	r, err := global.REDIS.Get(context.TODO(), key).Result()
	if err != nil {
		return err
	}
	if code != r {
		return errors.New("code not match, try again")
	}
	err = global.DB.Model(&dbTable.User{}).Where("UUID = ? ", u.UUID).Update("active", 1).Error
	if err != nil {
		return err
	}
	return global.REDIS.Del(context.TODO(), key).Err()
}

func (userService *UserService) Login(u *dbTable.User) (*dbTable.User, error) {
	var user dbTable.User
	err := global.DB.Where("username = ?", u.Username).Preload("Role").First(&user).Error
	if err == nil {
		if ok := utils.BcryptCheck(u.Password, user.Password); !ok {
			return nil, errors.New("wrong password")
		}
	}
	return &user, err
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

	u.Role = dbTable.Role{Name: "guest", ID: 3}
	u.Password = utils.BcryptHash(u.Password)
	u.UUID = uuid.NewV4()
	err = global.DB.Create(&u).Error
	return u, err
}

func (userService *UserService) EditUser(u dbTable.User) error {
	return global.DB.Model(&dbTable.User{}).Where("UUID = ? ", u.UUID).Updates(&u).Error
}

func (userService *UserService) DeleteUser(uuid string) error {
	return global.DB.Where("UUID = ?", uuid).Delete(&dbTable.User{}).Error
}

func (userService *UserService) ResetPassword(u dbTable.User, newPassword string, code string, uuid string) error {
	key := global.CONFIG.Email.Prefix + uuid
	r, err := global.REDIS.Get(context.TODO(), key).Result()
	if err != nil {
		return err
	}
	if code != r {
		return errors.New("code not match, try again")
	}
	global.REDIS.Del(context.TODO(), key).Err()
	newPassword = utils.BcryptHash(newPassword)

	return global.DB.Model(&dbTable.User{}).Where("id = ? ", u.ID).Update("password", newPassword).Error
}
