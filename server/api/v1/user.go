package apiV1

import (
	"github.com/ChocolateAceCream/blog/api/v1/shared/response"
	"github.com/ChocolateAceCream/blog/global"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserApi struct{}

// @Summary get user list
// @Description get user list
// @Tags Test
// @Accept json
// @Success 200 {object} response.Response{data=response.Paging{data=[]model.User},msg=string} "paged user list, includes page size, page number, total counts"
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
