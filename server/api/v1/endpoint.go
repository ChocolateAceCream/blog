/*
* @fileName endpoint.go
* @author Di Sheng
* @date 2023/02/12 21:15:51
* @description endpoint api
 */

package apiV1

import (
	"github.com/ChocolateAceCream/blog/global"
	"github.com/ChocolateAceCream/blog/model/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type EndpointApi struct{}

// @Summary get all endpoints
// @Description return all endpoints
// @Tags Endpoint
// @Accept json
// @Success 200 {object} response.Response{data=[]dbTable.Endpoint} "return all endpoints"
// @Router /api/v1/endpoint/getAllEndpoints [GET]
func (b *EndpointApi) GetAllEndpoints(c *gin.Context) {
	if r, err := endpointService.GetAllEndpoints(); err != nil {
		global.LOGGER.Error("fail to get all endpoints", zap.Error(err))
		response.FailWithMessage("fail to get all endpoints", c)
	} else {
		response.OkWithFullDetails(r, "success", c)
	}
	return
}
