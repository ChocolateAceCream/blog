package apiV1

import (
	"github.com/ChocolateAceCream/blog/global"
	"github.com/ChocolateAceCream/blog/model"
	"github.com/ChocolateAceCream/blog/model/dbTable"
	"github.com/ChocolateAceCream/blog/model/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserApi struct{}

// @Summary get user list
// @Description get user list
// @Tags User
// @Accept json
// @Success 200 {object} response.Response{data=response.Paging{data=[]dbTable.User},msg=string} "paged user list, includes page size, page number, total counts"
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

// @Tags User
// @Summary Register user
// @Description register user
// @Produce  application/json
// @Param data body model.Register true "username, password, email, role ID"
// @Success 200 {object} response.Response{data=dbTable.User,msg=string} "register user, return user info"
// @Router /v1/user/register [post]
func (b *UserApi) Register(c *gin.Context) {
	var r model.Register
	if err := c.ShouldBindJSON(&r); err != nil {
		global.LOGGER.Error("register user validation error", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	// var authorities []system.SysAuthority
	// for _, v := range r.AuthorityIds {
	// 	authorities = append(authorities, system.SysAuthority{
	// 		AuthorityId: v,
	// 	})
	// }
	payload := &dbTable.User{
		Username: r.Username,
		Password: r.Password,
		Email:    r.Email,
	}
	result, err := userService.RegisterUser(*payload)
	if err != nil {
		global.LOGGER.Error("Failed to Register!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithFullDetails(result, "User Register success", c)
	}
}

// @Tags User
// @Summary Edit user info
// @Description update user info
// @Produce  application/json
// @Param data body model.EditUser true "username, email, role ID,active, uuid"
// @Success 200 {object} response.Response{data=dbTable.User,msg=string} "edit user, return updated user info"
// @Router /v1/user/edit [put]
func (b *UserApi) EditUser(c *gin.Context) {
	var r model.EditUser
	if err := c.ShouldBindJSON(&r); err != nil {
		global.LOGGER.Error("edit user param validation error", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	//TODO: add auth check
	// var authorities []system.SysAuthority
	// for _, v := range r.AuthorityIds {
	// 	authorities = append(authorities, system.SysAuthority{
	// 		AuthorityId: v,
	// 	})
	// }
	payload := &dbTable.User{
		UUID:     r.UUID,
		Username: r.Username,
		Email:    r.Email,
		Active:   r.Active,
	}
	err := userService.EditUser(*payload)
	if err != nil {
		global.LOGGER.Error("Failed to edit user!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("Edit user success", c)
	}
}
