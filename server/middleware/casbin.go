package middleware

import (
	"net/http"
	"strconv"

	"github.com/ChocolateAceCream/blog/global"
	"github.com/ChocolateAceCream/blog/model/dbTable"
	"github.com/ChocolateAceCream/blog/model/response"
	"github.com/ChocolateAceCream/blog/service"
	"github.com/ChocolateAceCream/blog/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CasbinHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		currentUser, err := utils.GetValueFromSession[dbTable.User](c, "currentUser")
		if err != nil {
			global.LOGGER.Error("Fail to get current user", zap.Error(err))
			response.FailWithUnauthorized("Current Session Expired", c)
			c.Abort()
			return
		}
		path := c.Request.URL.Path
		// 获取请求方法
		action := c.Request.Method
		e := service.GetEnforcer()
		var roleService = new(service.RoleService)
		role, err := roleService.GetUserRole(currentUser)
		if err != nil {
			global.LOGGER.Error("Fail to get current user's Role", zap.Error(err))
			response.FailWithUnauthorized("Current Session Expired", c)
			c.Abort()
			return
		}
		id := strconv.Itoa(int(role.ID))
		if success, _ := e.Enforce(id, path, action); success {
			c.Next()
			return
		}
		response.FailWithFullDetails(http.StatusForbidden, "Access denied", c)
		c.Abort()
	}
}
