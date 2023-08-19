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

type ReplyApi struct{}

// @Tags Reply
// @Summary add new reply
// @Description add new reply
// @Accept json
// @Param data body request.AddReplyPayload true "replyContent, articleId "
// @Success 200 {object} response.Response{msg=string} "success"
// @Router /api/v1/reply/add [POST]
func (*ReplyApi) AddReply(c *gin.Context) {
	var payload request.AddReplyPayload

	if err := c.ShouldBindJSON(&payload); err != nil {
		global.LOGGER.Error("fail to bind reply data", zap.Error(err))
		response.FailWithMessage("fail to bind reply data, "+err.Error(), c)
		return
	}

	reply := dbTable.Reply{
		ReplyContent: payload.ReplyContent,
		CommentID:    payload.CommentID,
	}

	if payload.ParentReplyID != nil {
		reply.ParentReplyID = payload.ParentReplyID
	}

	currentUser, err := utils.GetValueFromSession[dbTable.User](c, "currentUser")
	if err != nil {
		global.LOGGER.Error("User not logged in", zap.Error(err))
		response.FailWithUnauthorized("User not logged in", c)
		return
	}
	reply.AuthorID = currentUser.ID

	if err := replyService.AddReply(reply); err != nil {
		global.LOGGER.Error("fail to add reply", zap.Error(err))
		response.FailWithMessage("fail to add reply, "+err.Error(), c)
	} else {
		response.OkWithMessage("add reply success", c)
	}
}

// @Tags      Reply
// @Summary   delete reply by id
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.FindById                true  "reply id"
// @Success   200   {object}  response.Response{msg=string}  "reply deleted "
// @Router 		/api/v1/reply/delete [delete]
func (a *ReplyApi) DeleteReply(c *gin.Context) {
	var reply request.FindById
	currentUser, err := utils.GetValueFromSession[dbTable.User](c, "currentUser")
	if err != nil {
		global.LOGGER.Error("User not logged in", zap.Error(err))
		response.FailWithUnauthorized("User not logged in", c)
		return
	}
	if err := c.ShouldBindJSON(&reply); err != nil {
		global.LOGGER.Error("delete reply params parsing error", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := replyService.DeleteReply(currentUser.ID, reply.ID); err != nil {
		global.LOGGER.Error("fail to delete reply", zap.Error(err))
		response.FailWithMessage("fail to delete reply", c)
	} else {
		response.OkWithMessage("delete reply success", c)
	}
}

// @Tags reply
// @Summary get reply list by cursorId
// @Description return reply list
// @Accept json
// @Param data query request.ReplyCursorListQuery true "get paged reply list by search query"
// @Success 200 {object} response.Response{data=response.Paging{list=[]dbTable.Reply,total=int}} "return all search result "
// @Router /api/v1/reply/list [GET]
func (b *ReplyApi) GetReplyList(c *gin.Context) {
	var query request.ReplyCursorListQuery
	if err := c.ShouldBind(&query); err != nil {
		global.LOGGER.Error("bind reply list query error", zap.Error(err))
		response.FailWithMessage("fail to bind query", c)
		return
	}

	currentUser, err := utils.GetValueFromSession[dbTable.User](c, "currentUser")
	if err != nil {
		global.LOGGER.Error("User not logged in", zap.Error(err))
		response.FailWithUnauthorized("User not logged in", c)
		return
	}

	if list, total, err := replyService.GetReplyList(query.Params, currentUser); err != nil {
		global.LOGGER.Error("fail to get reply list", zap.Error(err))
		response.FailWithMessage("fail to get reply list", c)
		return
	} else {
		response.OkWithFullDetails(response.Paging{
			List:  list,
			Total: total,
		}, "Success", c)
	}
}

// // @Tags Reply
// // @Summary like reply
// // @Description current user like reply or unlike reply
// // @Accept json
// // @Param data body request.LikeReplyPayload true "replyId, like "
// // @Success 200 {object} response.Response{msg=string} "success"
// // @Router /api/v1/reply/like [POST]
// func (*ReplyApi) LikeReply(c *gin.Context) {
// 	var payload request.LikeReplyPayload

// 	if err := c.ShouldBindJSON(&payload); err != nil {
// 		global.LOGGER.Error("fail to bind like reply data", zap.Error(err))
// 		response.FailWithMessage("fail to bind like reply data, "+err.Error(), c)
// 		return
// 	}

// 	currentUser, err := utils.GetValueFromSession[dbTable.User](c, "currentUser")
// 	if err != nil {
// 		global.LOGGER.Error("User not logged in", zap.Error(err))
// 		response.FailWithUnauthorized("User not logged in", c)
// 		return
// 	}
// 	payload.UserID = currentUser.ID
// 	if err := replyService.LikeReply(payload); err != nil {
// 		global.LOGGER.Error("failed", zap.Error(err))
// 		response.FailWithMessage("operation failed", c)
// 	} else {
// 		response.OkWithMessage("success", c)
// 	}
// }
