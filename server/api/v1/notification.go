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
	}
}

// @Tags notification
// @Summary get unread notification count
// @Description return unread notification count
// @Accept json
// @Success 200 {object} response.Response{data=int,msg=string} "Return current user's unread notification count"
// @Router /api/v1/notification/unreadCount [GET]
func (a *NotificationApi) GetUnreadCount(c *gin.Context) {
	currentUser, err := utils.GetValueFromSession[dbTable.User](c, "currentUser")
	if err != nil {
		global.LOGGER.Error("User not logged in", zap.Error(err))
		response.FailWithUnauthorized("User not logged in", c)
		return
	}
	if count, err := notificationService.GetUnreadCount(currentUser.ID); err != nil {
		global.LOGGER.Error("Fail to get unread notification count!", zap.Error(err))
		response.FailWithMessage("Fail to get unread notification count", c)
	} else {
		response.OkWithFullDetails(count, "success", c)
	}
}

// @Tags notification
// @Summary get notification list by cursorId
// @Description return notification list
// @Accept json
// @Param data query request.NotificationCursorListQuery true "get paged notification list by search query"
// @Success 200 {object} response.Response{data=response.Paging{list=[]dbTable.Notification,total=int}} "return all search result "
// @Router /api/v1/notification/list [GET]
func (a *NotificationApi) GetNotificationList(c *gin.Context) {
	var query request.NotificationCursorListQuery
	if err := c.ShouldBind(&query); err != nil {
		global.LOGGER.Error("bind notification list query error", zap.Error(err))
		response.FailWithMessage("fail to bind query", c)
		return
	}

	currentUser, err := utils.GetValueFromSession[dbTable.User](c, "currentUser")
	if err != nil {
		global.LOGGER.Error("User not logged in", zap.Error(err))
		response.FailWithUnauthorized("User not logged in", c)
		return
	}

	if list, total, err := commentService.GetNotificationList(query.Params, currentUser); err != nil {
		global.LOGGER.Error("fail to get notification list", zap.Error(err))
		response.FailWithMessage("fail to get notification list", c)
		return
	} else {
		response.OkWithFullDetails(response.Paging{
			List:  list,
			Total: total,
		}, "Success", c)
	}
}

// @Tags notification
// @Summary read notification
// @Description change notification status to read
// @Accept json
// @Param     data  body      request.FindById                true  "notification id"
// @Success 200 {object} response.Response{msg=string} "Change notification status to read"
// @Router /api/v1/notification/read [PATCH]
func (a *NotificationApi) ReadNotification(c *gin.Context) {
	var notification request.FindById
	currentUser, err := utils.GetValueFromSession[dbTable.User](c, "currentUser")
	if err != nil {
		global.LOGGER.Error("User not logged in", zap.Error(err))
		response.FailWithUnauthorized("User not logged in", c)
		return
	}
	if err := c.ShouldBindJSON(&notification); err != nil {
		global.LOGGER.Error("ReadNotification params parsing error", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := notificationService.ReadNotification(currentUser.ID, notification.ID); err != nil {
		global.LOGGER.Error("Fail to change notification status!", zap.Error(err))
		response.FailWithMessage("Fail to change notification status", c)
	} else {
		response.OkWithMessage("success", c)
	}
}

// @Tags      Notification
// @Summary   delete notification by id
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.FindById                true  "notification id"
// @Success   200   {object}  response.Response{msg=string}  "notification deleted "
// @Router 		/api/v1/notification/delete [delete]
func (a *NotificationApi) DeleteNotification(c *gin.Context) {
	var notification request.FindById
	currentUser, err := utils.GetValueFromSession[dbTable.User](c, "currentUser")
	if err != nil {
		global.LOGGER.Error("User not logged in", zap.Error(err))
		response.FailWithUnauthorized("User not logged in", c)
		return
	}
	if err := c.ShouldBindJSON(&notification); err != nil {
		global.LOGGER.Error("delete notification params parsing error", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := notificationService.DeleteNotification(currentUser.ID, notification.ID); err != nil {
		global.LOGGER.Error("fail to delete notification", zap.Error(err))
		response.FailWithMessage("fail to delete notification", c)
	} else {
		response.OkWithMessage("delete notification success", c)
	}
}
