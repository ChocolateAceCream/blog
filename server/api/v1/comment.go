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

type CommentApi struct{}

// @Tags Comment
// @Summary like comment
// @Description current user like comment or unlike comment
// @Accept json
// @Param data body request.LikeCommentPayload true "commentId, like "
// @Success 200 {object} response.Response{msg=string} "success"
// @Router /api/v1/comment/like [POST]
func (*CommentApi) LikeComment(c *gin.Context) {
	var payload request.LikeCommentPayload

	if err := c.ShouldBindJSON(&payload); err != nil {
		global.LOGGER.Error("fail to bind like comment data", zap.Error(err))
		response.FailWithMessage("fail to bind like comment data, "+err.Error(), c)
		return
	}

	currentUser, err := utils.GetValueFromSession[dbTable.User](c, "currentUser")
	if err != nil {
		global.LOGGER.Error("User not logged in", zap.Error(err))
		response.FailWithUnauthorized("User not logged in", c)
		return
	}
	payload.UserID = currentUser.ID
	if err := commentService.LikeComment(payload); err != nil {
		global.LOGGER.Error("failed", zap.Error(err))
		response.FailWithMessage("operation failed", c)
	} else {
		response.OkWithMessage("success", c)
	}
}

// @Tags Comment
// @Summary add new comment
// @Description add new comment
// @Accept json
// @Param data body request.AddCommentPayload true "commentContent, articleId "
// @Success 200 {object} response.Response{msg=string} "success"
// @Router /api/v1/comment/add [POST]
func (*CommentApi) AddComment(c *gin.Context) {
	var payload request.AddCommentPayload

	if err := c.ShouldBindJSON(&payload); err != nil {
		global.LOGGER.Error("fail to bind comment data", zap.Error(err))
		response.FailWithMessage("fail to bind comment data, "+err.Error(), c)
		return
	}

	currentUser, err := utils.GetValueFromSession[dbTable.User](c, "currentUser")
	if err != nil {
		global.LOGGER.Error("User not logged in", zap.Error(err))
		response.FailWithUnauthorized("User not logged in", c)
		return
	}
	comment := dbTable.Comment{
		AuthorID:       currentUser.ID,
		CommentContent: payload.CommentContent,
		ArticleID:      payload.ArticleID,
	}

	if err := commentService.AddComment(comment, payload.ArticleID); err != nil {
		global.LOGGER.Error("fail to add comment", zap.Error(err))
		response.FailWithMessage("fail to add comment, "+err.Error(), c)
	} else {
		response.OkWithMessage("add comment success", c)
	}
}

// @Tags comment
// @Summary get comment list by cursorId
// @Description return comment list
// @Accept json
// @Param data query request.CommentCursorListQuery true "get paged comment list by search query"
// @Success 200 {object} response.Response{data=response.Paging{list=[]dbTable.Comment,total=int}} "return all search result "
// @Router /api/v1/comment/list [GET]
func (b *CommentApi) GetCommentList(c *gin.Context) {
	var query request.CommentCursorListQuery
	if err := c.ShouldBind(&query); err != nil {
		global.LOGGER.Error("bind comment list query error", zap.Error(err))
		response.FailWithMessage("fail to bind query", c)
		return
	}

	currentUser, err := utils.GetValueFromSession[dbTable.User](c, "currentUser")
	if err != nil {
		global.LOGGER.Error("User not logged in", zap.Error(err))
		response.FailWithUnauthorized("User not logged in", c)
		return
	}

	if list, total, err := commentService.GetCommentList(query.Params, currentUser); err != nil {
		global.LOGGER.Error("fail to get comment list", zap.Error(err))
		response.FailWithMessage("fail to get comment list", c)
		return
	} else {
		response.OkWithFullDetails(response.Paging{
			List:  list,
			Total: total,
		}, "Success", c)
	}
}

// @Tags      Comment
// @Summary   delete comment by id
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.FindById                true  "comment id"
// @Success   200   {object}  response.Response{msg=string}  "comment deleted "
// @Router 		/api/v1/comment/delete [delete]
func (a *CommentApi) DeleteComment(c *gin.Context) {
	var comment request.FindById
	currentUser, err := utils.GetValueFromSession[dbTable.User](c, "currentUser")
	if err != nil {
		global.LOGGER.Error("User not logged in", zap.Error(err))
		response.FailWithUnauthorized("User not logged in", c)
		return
	}
	if err := c.ShouldBindJSON(&comment); err != nil {
		global.LOGGER.Error("delete comment params parsing error", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := commentService.DeleteComment(currentUser.ID, comment.ID); err != nil {
		global.LOGGER.Error("fail to delete comment", zap.Error(err))
		response.FailWithMessage("fail to delete comment", c)
	} else {
		response.OkWithMessage("delete comment success", c)
	}
}
