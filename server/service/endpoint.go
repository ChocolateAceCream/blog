/*
* @fileName endpoint.go
* @author Di Sheng
* @date 2023/02/12 21:16:12
* @description endpoint service
 */

package service

import (
	"github.com/ChocolateAceCream/blog/global"
	"github.com/ChocolateAceCream/blog/model/dbTable"
)

type EndpointService struct{}

func (es *EndpointService) GetAllEndpoints() (endpoints []dbTable.Endpoint, err error) {
	err = global.DB.Find(&endpoints).Error
	return
}
