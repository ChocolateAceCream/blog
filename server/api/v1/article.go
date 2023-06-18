package apiV1

import (
	"github.com/ChocolateAceCream/blog/global"
	"github.com/ChocolateAceCream/blog/model/request"
	"github.com/ChocolateAceCream/blog/model/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ArticleApi struct{}

// @Summary preview article
// @Description get article md files and other info
// @Tags Article
// @Accept json
// @Param data query request.PreviewArticle true "get article md file by article id"
// @Success 200 {object} response.Response{data=response.ArticleInfo,msg=string} "return article info & md file"
// @Router /api/v1/article/preview [get]
func (b *ArticleApi) PreviewArticle(c *gin.Context) {
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
