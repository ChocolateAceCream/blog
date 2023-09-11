package apiV1

import (
	"github.com/ChocolateAceCream/blog/global"
	"github.com/ChocolateAceCream/blog/model/dbTable"
	"github.com/ChocolateAceCream/blog/model/response"
	"github.com/ChocolateAceCream/blog/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type NotificationApi struct{}

func (a *NotificationApi) WSHandler(c *gin.Context) {
	currentUser, err := utils.GetValueFromSession[dbTable.User](c, "currentUser")
	if err != nil {
		global.LOGGER.Error("User not logged in", zap.Error(err))
		response.FailWithUnauthorized("User not logged in", c)
		return
	}

	if err := notificationService.WSHandler(c, currentUser.ID); err != nil {
		global.LOGGER.Error("fail to handle websocket", zap.Error(err))
		response.FailWithMessage("fail to handle websocket, "+err.Error(), c)
	}
}
