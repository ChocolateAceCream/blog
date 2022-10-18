package apiV1

import (
	"github.com/ChocolateAceCream/blog/api/v1/shared/response"
	"github.com/ChocolateAceCream/blog/global"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserApi struct{}

// @Summary 测试SayHello
// @Description 向你说Hello
// @Tags 测试
// @Accept json
// @Success 200 {object} response.Paging"
// @Failure 200 {string} response.Response"
// @Router /v1/user/userList [get]
func (b *UserApi) GetUserList(c *gin.Context) {
	if list, total, err := userService.GetUserInfoList(); err != nil {
		global.LOGGER.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithFullDetails(response.Paging{
			List:  list,
			Total: total,
		}, "Success", c)
	}
}
