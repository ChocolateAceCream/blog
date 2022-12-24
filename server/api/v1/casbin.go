package apiV1

import (
	"github.com/ChocolateAceCream/blog/global"
	"github.com/ChocolateAceCream/blog/model/request"
	"github.com/ChocolateAceCream/blog/model/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CasbinApi struct{}

// @Tags Casbin
// @Summary Update role privilege
// @accept application/json
// @Produce application/json
// @Param data body request.UpdateCasbin true "update privilege based on role id"
// @Success 200 {object} response.Response{msg=string} "update privilege based on role id"
// @Router /api/v1/casbin/update [post]
func (cas *CasbinApi) UpdateCasbin(c *gin.Context) {
	var r request.UpdateCasbin
	if err := c.ShouldBindJSON(&r); err != nil {
		global.LOGGER.Error("update casbin error", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := casbinService.Update(r.RoleId, r.Endpoints); err != nil {
		global.LOGGER.Error("Fail to update!", zap.Error(err))
		response.FailWithMessage("Fail to update casbin", c)
	} else {
		response.OkWithMessage("Update success", c)
	}
}
