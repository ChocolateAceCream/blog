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
			response.FailWithMessage(err.Error(), c)
			c.Abort()
			return
		}
		path := c.Request.URL.Path
		// 获取请求方法
		action := c.Request.Method
		e := service.GetEnforcer()
		for _, role := range currentUser.UserRoles {
			id := strconv.Itoa(int(role.RoleId))
			if success, _ := e.Enforce(id, path, action); success {
				c.Next()
				return
			}
		}
		response.FailWithFullDetails(http.StatusUnauthorized, "Access denied", c)
		c.Abort()
	}
}
