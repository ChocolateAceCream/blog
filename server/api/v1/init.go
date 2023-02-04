package apiV1

import (
	"github.com/ChocolateAceCream/blog/global"
	"github.com/ChocolateAceCream/blog/model/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type InitApi struct{}

// Init db data
// @Tags Init
// @Summary initialize data
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "return init result msg"
// @Router /api/public/initDB [post]
func (i *InitApi) InitDB(c *gin.Context) {
	if err := initService.InitDB(); err != nil {
		global.LOGGER.Error("Fail to init db data!", zap.Error(err))
		response.FailWithMessage("Fail to init db data", c)
		return
	}
	response.OkWithMessage("create role success", c)
}
