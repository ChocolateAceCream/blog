package apiV1

import (
	"github.com/ChocolateAceCream/blog/global"
	"github.com/ChocolateAceCream/blog/model/dbTable"
	"github.com/ChocolateAceCream/blog/model/request"
	"github.com/ChocolateAceCream/blog/model/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type RoleApi struct{}

// @Tags Role
// @Summary Create Role
// @accept application/json
// @Produce application/json
// @Param data body dbTable.Role true "name, parentId"
// @Success 200 {object} response.Response{data=dbTable.Role,msg=string} "Create role, return"
// @Router /api/v1/role/create [post]
func (a *RoleApi) CreateRole(c *gin.Context) {
	var r dbTable.Role

	if err := c.ShouldBindJSON(&r); err != nil {
		response.FailWithMessage("Failed to parse query", c)
	}

	if r, err := roleService.CreateRole(r); err != nil {
		global.LOGGER.Error("Fail to create role!", zap.Error(err))
		response.FailWithMessage("Fail to create role! "+err.Error(), c)
	} else {
		_ = casbinService.Update(r.RoleId, request.DefaultCasbin())
		response.OkWithFullDetails(r, "create role success", c)
	}
}
