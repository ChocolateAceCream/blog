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

// @Tags article
// @Summary get article list by cursorId
// @Description return article list
// @Accept json
// @Param data query request.ArticleCursorListQuery true "get paged article list by search query"
// @Success 200 {object} response.Response{data=response.Paging{list=[]dbTable.Article,total=int}} "return all search result "
// @Router /api/v1/article/list [GET]
func (b *ArticleApi) GetArticleList(c *gin.Context) {
	var query request.ArticleCursorListQuery
	if err := c.ShouldBind(&query); err != nil {
		global.LOGGER.Error("bind article list query error", zap.Error(err))
		response.FailWithMessage("fail to bind query", c)
		return
	}
	if list, total, err := articleService.GetArticleList(query.Params); err != nil {
		global.LOGGER.Error("fail to get article list", zap.Error(err))
		response.FailWithMessage("fail to get article list", c)
		return
	} else {
		response.OkWithFullDetails(response.Paging{
			List:  list,
			Total: total,
		}, "Success", c)
	}
}

// @Tags Article
// @Summary get article search list
// @Description return article search result in paged list, only applied to articles with current user as author
// @Accept json
// @Param data query request.ArticleSearchQuery true "get paged article list by search query"
// @Success 200 {object} response.Response{data=response.Paging{list=[]dbTable.Article,total=int}} "return all search result "
// @Router /api/v1/article/list [GET]
func (b *ArticleApi) GetArticleSearchList(c *gin.Context) {
	currentUser, err := utils.GetValueFromSession[dbTable.User](c, "currentUser")
	if err != nil {
		global.LOGGER.Error("User not logged in", zap.Error(err))
		response.FailWithUnauthorized("User not logged in", c)
		return
	}
	var query request.ArticleSearchQuery
	if err := c.ShouldBind(&query); err != nil {
		global.LOGGER.Error("bind article list query error", zap.Error(err))
		response.FailWithMessage("fail to bind query", c)
		return
	}
	if list, total, err := articleService.GetArticleSearchList(currentUser.ID, query.Params); err != nil {
		global.LOGGER.Error("fail to get article list", zap.Error(err))
		response.FailWithMessage("fail to get article list", c)
		return
	} else {
		response.OkWithFullDetails(response.Paging{
			List:  list,
			Total: total,
		}, "Success", c)
	}
}

// @Tags      Article
// @Summary   delete article by id
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.FindById                true  "article id"
// @Success   200   {object}  response.Response{msg=string}  "article deleted "
// @Router 		/api/v1/article/delete [delete]
func (a *ArticleApi) DeleteArticle(c *gin.Context) {
	var article request.FindByIds
	currentUser, err := utils.GetValueFromSession[dbTable.User](c, "currentUser")
	if err != nil {
		global.LOGGER.Error("User not logged in", zap.Error(err))
		response.FailWithUnauthorized("User not logged in", c)
		return
	}
	if err := c.ShouldBindJSON(&article); err != nil {
		global.LOGGER.Error("delete article params parsing error", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := articleService.DeleteArticle(currentUser.ID, article.ID); err != nil {
		global.LOGGER.Error("fail to delete article", zap.Error(err))
		response.FailWithMessage("fail to delete article", c)
	} else {
		response.OkWithMessage("delete article success", c)
	}
}
