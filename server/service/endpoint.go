/*
* @fileName endpoint.go
* @author Di Sheng
* @date 2023/02/12 21:16:12
* @description endpoint service
 */

package service

import (
	"errors"

	"github.com/ChocolateAceCream/blog/global"
	"github.com/ChocolateAceCream/blog/model/dbTable"
	"gorm.io/gorm"
)

type EndpointService struct{}

func (es *EndpointService) GetAllEndpoints() (endpoints []dbTable.Endpoint, err error) {
	err = global.DB.Find(&endpoints).Error
	return
}

func (es *EndpointService) NewEndpoint(endpoint dbTable.Endpoint) error {
	db := global.DB
	err := db.Where("method = ? and path = ?", endpoint.Method, endpoint.Path).First(&dbTable.Endpoint{}).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("endpoint already existed")
	}
	return db.Create(&endpoint).Error
}
