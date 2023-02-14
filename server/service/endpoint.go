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
	"github.com/ChocolateAceCream/blog/model/request"
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

func (es *EndpointService) GetEndpointList(query request.EndpointSearchQuery) (endpointList []dbTable.Endpoint, total int64, err error) {
	db := global.DB.Model(&dbTable.Endpoint{})
	if query.Method != "" {
		db.Where("method LIKE ? ", "%"+query.Method+"%")
	}
	if query.Name != "" {
		db.Where("name LIKE ? ", "%"+query.Name+"%")
	}
	if query.Path != "" {
		db.Where("path LIKE ? ", "%"+query.Path+"%")
	}
	if query.Type != 0 {
		db.Where("type = ? ", query.Type)
	}
	if query.PID != 0 {
		db.Where("pid = ? ", query.PID)
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
		orderMap := make(map[string]bool, 4)
		orderMap["name"] = true
		orderMap["pid"] = true
		orderMap["path"] = true
		orderMap["id"] = true
		if orderMap[query.OrderBy] {
			if query.Desc {
				orderStr = query.OrderBy + " desc"
			} else {
				orderStr = query.OrderBy
			}
		}
		err = db.Order(orderStr).Find(&endpointList).Error
	} else {
		err = db.Order("id").Find(&endpointList).Offset(-1).Limit(-1).Error
	}
	return
}
