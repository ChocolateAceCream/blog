/*
* @fileName endpoint.go
* @author Di Sheng
* @date 2023/02/12 21:15:51
* @description endpoint api
 */

package apiV1

import (
	"github.com/ChocolateAceCream/blog/global"
	"github.com/ChocolateAceCream/blog/model/dbTable"
	"github.com/ChocolateAceCream/blog/model/request"
	"github.com/ChocolateAceCream/blog/model/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type EndpointApi struct{}

// @Summary get endpoint list
// @Description return endpoint list
// @Tags Endpoint
// @Accept json
// @Param data query request.EndpointSearchQuery true "get paged endpoint list by search query"
// @Success 200 {object} response.Response{data=[]dbTable.Endpoint} "return all search result "
// @Router /api/v1/endpoint/list [GET]
func (b *EndpointApi) GetEndpointList(c *gin.Context) {
	var query request.EndpointSearchQuery
	if err := c.ShouldBind(&query); err != nil {
		global.LOGGER.Error("bind endpoint list query error", zap.Error(err))
		response.FailWithMessage("fail to bind query", c)
		return
	}
	if list, total, err := endpointService.GetEndpointList(query.Params); err != nil {
		global.LOGGER.Error("fail to get endpoint list", zap.Error(err))
		response.FailWithMessage("fail to get endpoint list", c)
		return
	} else {
		response.OkWithFullDetails(response.Paging{
			List:  list,
			Total: total,
		}, "Success", c)
	}
}

// @Summary add new endpoint
// @Description add new endpoint
// @Tags Endpoint
// @Accept json
// @Param data body dbTable.Endpoint true "Type, Name, PID, Method, Description, Path "
// @Success 200 {object} response.Response{msg=string} "success"
// @Router /api/v1/endpoint/add [POST]
func (b *EndpointApi) AddEndpoint(c *gin.Context) {
	var r dbTable.Endpoint
	if err := c.ShouldBindJSON(&r); err != nil {
		global.LOGGER.Error("fail to bind endpoint data", zap.Error(err))
		response.FailWithMessage("fail to bind endpoint data, err: "+err.Error(), c)
		return
	}

	if err := endpointService.AddEndpoint(r); err != nil {
		global.LOGGER.Error("fail to add endpoint", zap.Error(err))
		response.FailWithMessage("fail to add endpoint, err: "+err.Error(), c)
	} else {
		response.OkWithMessage("add endpoint success", c)
	}
}
