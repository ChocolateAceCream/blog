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

// @Tags Endpoint
// @Summary get endpoint list
// @Description return endpoint list
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

// @Tags Endpoint
// @Summary add new endpoint
// @Description add new endpoint
// @Accept json
// @Param data body dbTable.Endpoint true "Group, Name, Method, Description, Path "
// @Success 200 {object} response.Response{msg=string} "success"
// @Router /api/v1/endpoint/add [POST]
func (b *EndpointApi) AddEndpoint(c *gin.Context) {
	var r dbTable.Endpoint
	if err := c.ShouldBindJSON(&r); err != nil {
		global.LOGGER.Error("fail to bind endpoint data", zap.Error(err))
		response.FailWithMessage("fail to bind endpoint data, "+err.Error(), c)
		return
	}

	if err := endpointService.AddEndpoint(r); err != nil {
		global.LOGGER.Error("fail to add endpoint", zap.Error(err))
		response.FailWithMessage("fail to add endpoint, "+err.Error(), c)
	} else {
		response.OkWithMessage("add endpoint success", c)
	}
}

// @Tags      Endpoint
// @Summary   edit endpoint
// @accept    application/json
// @Produce   application/json
// @Param     data  body      dbTable.Endpoint             true  "Group, Name, Method, Description, Path"
// @Success   200   {object}  response.Response{msg=string}  "success "
// @Router 		/api/v1/endpoint/edit [put]
func (a *EndpointApi) EditEndpoint(c *gin.Context) {
	var endpoint dbTable.Endpoint
	if err := c.ShouldBindJSON(&endpoint); err != nil {
		global.LOGGER.Error("edit endpoint params parsing error", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := endpointService.EditEndpoint(endpoint); err != nil {
		global.LOGGER.Error("fail to edit endpoint", zap.Error(err))
		response.FailWithMessage("fail to edit endpoint", c)
	} else {
		response.OkWithMessage("Edit endpoint success", c)
	}
}

// @Tags      Endpoint
// @Summary   delete endpoint by id
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.FindById                true  "endpoint id"
// @Success   200   {object}  response.Response{msg=string}  "endpoint deleted "
// @Router 		/api/v1/endpoint/delete [delete]
func (a *EndpointApi) DeleteEndpoint(c *gin.Context) {
	var endpoint request.FindByIds
	if err := c.ShouldBindJSON(&endpoint); err != nil {
		global.LOGGER.Error("delete endpoint params parsing error", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := endpointService.DeleteEndpoint(endpoint.ID); err != nil {
		global.LOGGER.Error("fail to delete endpoint", zap.Error(err))
		response.FailWithMessage("fail to delete endpoint", c)
	} else {
		response.OkWithMessage("delete endpoint success", c)
	}
}
