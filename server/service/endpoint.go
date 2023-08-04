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

func (es *EndpointService) AddEndpoint(endpoint dbTable.Endpoint) error {
	db := global.DB
	err := db.Where("method = ? and path = ?", endpoint.Method, endpoint.Path).First(&dbTable.Endpoint{}).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("endpoint already existed")
	}
	return db.Create(&endpoint).Error
}

func (es *EndpointService) EditEndpoint(endpoint dbTable.Endpoint) error {
	return global.DB.Model(&dbTable.Endpoint{}).Where("ID = ? ", endpoint.ID).Updates(&endpoint).Error
}

func (es *EndpointService) GetEndpointList(query request.EndpointSearchParam) (endpointList []dbTable.Endpoint, total int64, err error) {
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
	if query.GroupName != "" {
		db.Where("group_name LIKE ? ", "%"+query.GroupName+"%")
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
		orderMap := make(map[string]bool, 5)
		orderMap["name"] = true
		orderMap["method"] = true
		orderMap["path"] = true
		orderMap["id"] = true
		orderMap["group_name"] = true
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

func (endpointService *EndpointService) DeleteEndpoint(ids []int) (err error) {
	//TODO: test delete associated role-endpoint relations
	endpoints := []dbTable.Endpoint{}
	if err = global.DB.Find(&endpoints, "id in ?", ids).Delete(&endpoints).Error; err != nil {
		return err
	}
	for _, es := range endpoints {
		success := CasbinServiceInstance.ClearCasbin(1, es.Path, es.Method) // index of path in casbin rule is 1
		if !success {
			return errors.New("cannot clear casbin rule: " + es.Path + ":" + es.Method)
		}
	}
	return nil
}
