package apiV1

import (
	"github.com/ChocolateAceCream/blog/global"
	"github.com/ChocolateAceCream/blog/model/dbTable"
	"github.com/ChocolateAceCream/blog/model/request"
	"github.com/ChocolateAceCream/blog/model/response"
	"github.com/ChocolateAceCream/blog/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type MenuApi struct{}

// @Tags Menu
// @Summary add menu
// @accept application/json
// @Produce application/json
// @Param data body dbTable.Menu true "route path, pid, route name, component"
// @Success 200 {object} response.Response{msg=string} "新增菜单"
// @Router /api/v1/menu/add [post]
func (a *MenuApi) AddMenu(c *gin.Context) {
	var menu dbTable.Menu
	if err := c.ShouldBindJSON(&menu); err != nil {
		global.LOGGER.Error("add menu params parsing error", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := menuService.AddMenu(menu); err != nil {
		global.LOGGER.Error("fail to add menu", zap.Error(err))

		response.FailWithMessage("fail to add menu", c)
	} else {
		response.OkWithMessage("Add menu success", c)
	}
}

// @Tags Menu
// @Summary get current user's menus
// @Produce  application/json
// @Success 200 {object} response.Response{data=[]dbTable.Menu,msg=string} "Return current user's menu list"
// @Router /api/v1/menu/currentUserMenu [get]
func (a *MenuApi) GetCurrentUserMenu(c *gin.Context) {
	currentUser, err := utils.GetValueFromSession[dbTable.User](c, "currentUser")
	if err != nil {
		global.LOGGER.Error("User not logged in", zap.Error(err))
		response.FailWithUnauthorized("User not logged in", c)
		return
	}
	role, err := roleService.GetUserRole(currentUser)
	if err != nil {
		global.LOGGER.Error("Fail to get user role!", zap.Error(err))
		response.FailWithMessage("Fail to get user role", c)
		return
	}
	if menus, err := menuService.GetRoleMenus(role.ID); err != nil {
		global.LOGGER.Error("Fail to get menus!", zap.Error(err))
		response.FailWithMessage("Fail to get menus", c)
	} else {
		response.OkWithFullDetails(menus, "success", c)
	}
}

// @Tags Menu
// @Summary get menus by role id
// @Produce  application/json
// @Param     data  body      request.FindById                true  "role id"
// @Success 200 {object} response.Response{data=[]dbTable.Menu,msg=string} "Return role's menu list"
// @Router /api/v1/menu/getRoleMenuTree [post]
func (a *MenuApi) GetRoleMenuTree(c *gin.Context) {
	var role request.FindById
	if err := c.ShouldBindJSON(&role); err != nil {
		global.LOGGER.Error("Get role menu params parsing error", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	roleMenus, err := menuService.GetRoleMenus(uint(role.ID))
	if err != nil {
		global.LOGGER.Error("Fail to get role menus!", zap.Error(err))
		response.FailWithMessage("Fail to get role menus", c)
		return
	}
	menuList, err := menuService.GetMenuList()
	if err != nil {
		global.LOGGER.Error("Fail to get all menus!", zap.Error(err))
		response.FailWithMessage("Fail to get all menus", c)
		return
	}
	r := response.RoleMenuTree{
		MenuList:  menuList,
		RoleMenus: roleMenus,
	}
	response.OkWithFullDetails(r, "success", c)
}

// @Tags Menu
// @Summary get all menus
// @Produce  application/json
// @Success 200 {object} response.Response{data=[]dbTable.Menu,msg=string} "Return all menus"
// @Router /api/v1/menu/list [get]
func (a *MenuApi) GetMenuList(c *gin.Context) {
	if menus, err := menuService.GetMenuList(); err != nil {
		global.LOGGER.Error("Fail to get menus!", zap.Error(err))
		response.FailWithMessage("Fail to get menus", c)
	} else {
		response.OkWithFullDetails(menus, "success", c)
	}
}

// @Tags      Menu
// @Summary   delete menu by id
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.FindById                true  "menu id"
// @Success   200   {object}  response.Response{msg=string}  "menu deleted "
// @Router 		/api/v1/menu/delete [delete]
func (a *MenuApi) DeleteMenu(c *gin.Context) {
	var menu request.FindByIds
	if err := c.ShouldBindJSON(&menu); err != nil {
		global.LOGGER.Error("delete menu params parsing error", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := menuService.DeleteMenu(menu.ID); err != nil {
		global.LOGGER.Error("fail to delete menu", zap.Error(err))

		response.FailWithMessage("fail to delete menu", c)
	} else {
		response.OkWithMessage("delete menu success", c)
	}
}

// @Tags      Menu
// @Summary   edit menu
// @accept    application/json
// @Produce   application/json
// @Param     data  body      dbTable.Menu             true  "route, path, title, icon"
// @Success   200   {object}  response.Response{msg=string}  "menu edit success "
// @Router 		/api/v1/menu/edit [put]
func (a *MenuApi) EditMenu(c *gin.Context) {
	var menu dbTable.Menu
	if err := c.ShouldBindJSON(&menu); err != nil {
		global.LOGGER.Error("edit menu params parsing error", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := menuService.EditMenu(menu); err != nil {
		global.LOGGER.Error("fail to edit menu", zap.Error(err))
		response.FailWithMessage("fail to edit menu", c)
	} else {
		response.OkWithMessage("Edit menu success", c)
	}
}

// @Tags      Menu
// @Summary   edit menu
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.AssignRoleMenus           true  "id, menus"
// @Success   200   {object}  response.Response{msg=string}  "assign role menus success "
// @Router 		/api/v1/menu/assignRoleMenus [post]
func (a *MenuApi) AssignRoleMenus(c *gin.Context) {
	var payload request.AssignRoleMenus
	if err := c.ShouldBindJSON(&payload); err != nil {
		global.LOGGER.Error("assign role menus params parsing error", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := menuService.AssignRoleMenus(payload.Menus, uint(payload.ID)); err != nil {
		global.LOGGER.Error("fail to assign role menu", zap.Error(err))
		response.FailWithMessage("fail to assign role menus", c)
	} else {
		response.OkWithMessage("assign role menus success", c)
	}
}
