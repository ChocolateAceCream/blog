package apiV1

import (
	"fmt"

	"github.com/ChocolateAceCream/blog/global"
	"github.com/ChocolateAceCream/blog/model/dbTable"
	"github.com/ChocolateAceCream/blog/model/request"
	"github.com/ChocolateAceCream/blog/model/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type RoleApi struct{}

// @Tags Role
// @Summary Edit Role
// @accept application/json
// @Produce application/json
// @Param data body request.EditRole true "name, pid, id"
// @Success 200 {object} response.Response{data=dbTable.Role,msg=string} "edit role, return"
// @Router /api/v1/role/edit [put]
func (a *RoleApi) EditRole(c *gin.Context) {
	var r request.EditRole
	if err := c.ShouldBindJSON(&r); err != nil {
		response.FailWithMessage("Failed to parse query", c)
		return
	}
	if err := roleService.EditRole(dbTable.Role{Name: r.Name, ID: r.ID}); err != nil {
		global.LOGGER.Error("fail to edit role", zap.Error(err))
		response.FailWithMessage("fail to edit role", c)
	} else {
		response.OkWithMessage("Edit role success", c)
	}
}

// @Tags Role
// @Summary Create Role
// @accept application/json
// @Produce application/json
// @Param data body request.AddRole true "name, pid"
// @Success 200 {object} response.Response{data=dbTable.Role,msg=string} "Create role, return"
// @Router /api/v1/role/add [post]
func (a *RoleApi) AddRole(c *gin.Context) {
	var r request.AddRole

	if err := c.ShouldBindJSON(&r); err != nil {
		fmt.Println("------parse---", err)
		response.FailWithMessage("Failed to parse query", c)
		return
	}
	role, err := roleService.AddRole(dbTable.Role{Pid: *r.Pid, Name: r.Name})
	if err != nil {
		global.LOGGER.Error("Fail to add role!", zap.Error(err))
		response.FailWithMessage("Fail to add role! "+err.Error(), c)
	} else {
		if err := menuService.AssignRoleMenus(request.DefaultMenus(), role.ID); err != nil {
			global.LOGGER.Error("Fail to assign role menus!", zap.Error(err))
		}
		if err := casbinService.Update(role.ID, request.DefaultCasbin()); err != nil {
			global.LOGGER.Error("Fail to update role casbin!", zap.Error(err))
		}
		response.OkWithFullDetails(r, "add role success", c)
	}
}

// @Tags Role
// @Summary Get Role List
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=[]dbTable.Role,msg=string} "return all roles"
// @Router /api/v1/role/list [get]
func (a *RoleApi) GetRoleList(c *gin.Context) {
	if roles, err := roleService.GetRoleList(); err != nil {
		global.LOGGER.Error("Fail to get roles!", zap.Error(err))
		response.FailWithMessage("Fail to get roles", c)
	} else {
		response.OkWithFullDetails(roles, "success", c)
	}
}

// @Tags      Role
// @Summary   delete role by id
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.FindById                true  "role id"
// @Success   200   {object}  response.Response{msg=string}  "role deleted "
// @Router 		/api/v1/role/delete [delete]
func (a *RoleApi) DeleteRole(c *gin.Context) {
	var role request.FindByIds
	if err := c.ShouldBindJSON(&role); err != nil {
		global.LOGGER.Error("delete role params parsing error", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := roleService.DeleteRole(role.ID); err != nil {
		global.LOGGER.Error("fail to delete role", zap.Error(err))
		response.FailWithMessage("fail to delete role", c)
	} else {
		response.OkWithMessage("delete role success", c)
	}
}
