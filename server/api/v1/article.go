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

type ArticleApi struct{}

// @Summary preview article
// @Description get article md files and other info
// @Tags Article
// @Accept json
// @Param data query request.PreviewArticle true "get article md file by article id"
// @Success 200 {object} response.Response{data=request.PreviewArticle,msg=string} "return article info & md file"
// @Router /api/v1/article/preview [get]
func (*ArticleApi) PreviewArticle(c *gin.Context) {
	var r request.PreviewArticle
	if err := c.ShouldBind(&r); err != nil {
		global.LOGGER.Error("preview article validation error", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	if article, err := articleService.GetArticleInfo(r.Params.ID); err != nil {
		global.LOGGER.Error("Fail to get article info!", zap.Error(err))
		response.FailWithMessage("Fail to get article info", c)
	} else {
		response.OkWithFullDetails(article, "Success", c)
	}
}

// @Tags Article
// @Summary add new article
// @Description add new article
// @Accept json
// @Success 200 {object} response.Response{msg=string} "success"
// @Router /api/v1/article/add [POST]
func (*ArticleApi) AddArticle(c *gin.Context) {
	currentUser, err := utils.GetValueFromSession[dbTable.User](c, "currentUser")
	if err != nil {
		global.LOGGER.Error("User not logged in", zap.Error(err))
		response.FailWithUnauthorized("User not logged in", c)
		return
	}

	payload := dbTable.Article{
		Author: currentUser,
	}

	if r, err := articleService.AddArticle(payload); err != nil {
		global.LOGGER.Error("fail to add article", zap.Error(err))
		response.FailWithMessage("fail to add article, "+err.Error(), c)
	} else {
		response.OkWithFullDetails(r, "add article success", c)
	}
}

// @Tags Article
// @Summary edit article
// @Description edit article
// @Accept json
// @Param data body dbTable.Article true "Title, Content,Abstract"
// @Success 200 {object} response.Response{msg=string} "success"
// @Router /api/v1/article/edit [PUT]
func (*ArticleApi) EditArticle(c *gin.Context) {
	var a dbTable.Article
	if err := c.ShouldBindJSON(&a); err != nil {
		global.LOGGER.Error("fail to bind article data", zap.Error(err))
		response.FailWithMessage("fail to bind article data, "+err.Error(), c)
		return
	}

	currentUser, err := utils.GetValueFromSession[dbTable.User](c, "currentUser")
	if err != nil {
		global.LOGGER.Error("User not logged in", zap.Error(err))
		response.FailWithUnauthorized("User not logged in", c)
		return
	}
	hasPermission := articleService.HasPermission(currentUser.ID, a.ID)
	if !hasPermission {
		response.FailWithUnauthorized("not allowed", c)
		return
	}
	if err := articleService.EditArticle(a); err != nil {
		global.LOGGER.Error("fail to edit article", zap.Error(err))
		response.FailWithMessage("fail to edit article, "+err.Error(), c)
	} else {
		response.OkWithMessage("edit article success", c)
	}
}
