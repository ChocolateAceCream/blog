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

type UserApi struct{}

// @Summary get user list
// @Description get user list
// @Tags User
// @Accept json
// @Param data query request.UserSearchQuery true "get paged user list by search query"
// @Success 200 {object} response.Response{data=response.Paging{data=[]dbTable.User},msg=string} "paged user list, includes page size, page number, total counts"
// @Router /api/v1/user/userList [get]
func (b *UserApi) GetUserList(c *gin.Context) {
	// session := middleware.GetSession(c)
	// session.Set("asdf", 123)
	// a, _ := session.Get("asdf")
	// fmt.Println("key from session is: ", a)
	// session.RemoveKey("asdf")
	// a, _ = session.Get("asdf")
	// fmt.Println("after remove key from session, a is : ", a)

	var query request.UserSearchQuery
	// when using ShouldBindQuery, remember to comment form:"pageNumber" in data struct
	err := c.ShouldBindQuery(&query)
	if err != nil {
		response.FailWithMessage("Failed to parse query", c)
	}

	if list, total, err := userService.GetUserInfoList(query); err != nil {
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
// @Param data body request.RegisterUser true "username, password, email,captcha, role IDs"
// @Success 200 {object} response.Response{data=dbTable.User,msg=string} "register user, return user info"
// @Router /api/public/auth/register [post]
func (b *UserApi) RegisterUser(c *gin.Context) {
	var r request.RegisterUser
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
	payload := dbTable.User{
		Username: r.Username,
		Password: r.Password,
		Email:    r.Email,
		Active:   2, // inactive
	}

	u, err := userService.RegisterUser(payload)
	if err != nil {
		global.LOGGER.Error("Failed to Register!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithFullDetails(u, "User Register success", c)
		session := utils.GetSession(c)
		session.Set("currentUser", u)
	}
}

// @Tags User
// @Summary Active user
// @Description Active user with verification code from email
// @Produce  application/json
// @Param data body request.ActiveUser true "code"
// @Success 200 {object} response.Response{data=dbTable.User,msg=string} "return activated user info"
// @Router /api/v1/user/active [post]
func (b *UserApi) ActiveUser(c *gin.Context) {
	var r request.ActiveUser
	if err := c.ShouldBindJSON(&r); err != nil {
		global.LOGGER.Error("active user validation error", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	currentUser, err := utils.GetValueFromSession[dbTable.User](c, "currentUser")
	if err != nil {
		global.LOGGER.Error("active user validation error", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = userService.ActiveUser(currentUser, r.Code)
	if err != nil {
		global.LOGGER.Error("Failed to active user!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		currentUser.Active = 1
		response.OkWithFullDetails(currentUser, "User activated", c)
	}
}

// @Tags User
// @Summary Edit user info
// @Description update user info
// @Produce  application/json
// @Param data body request.EditUser true "username, email, role ID,active, uuid"
// @Success 200 {object} response.Response{data=dbTable.User,msg=string} "edit user, return updated user info"
// @Router /api/v1/user/edit [put]
func (b *UserApi) EditUser(c *gin.Context) {
	var r request.EditUser
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

// @Tags User
// @Summary delete user
// @accept application/json
// @Produce application/json
// @Param data body request.DeleteUser true "user uuid"
// @Success 200 {object} response.Response{msg=string} "user deleted"
// @Router /api/v1/user/delete [delete]
func (b *UserApi) DeleteUser(c *gin.Context) {
	var user request.DeleteUser
	_ = c.ShouldBindJSON(&user)

	// TODO: validate if target user is current user
	// jwtId := utils.GetUserID(c)
	// if jwtId == uint(reqId.ID) {
	// 	response.FailWithMessage("cannot delete yourself", c)
	// 	return
	// }
	if err := userService.DeleteUser(user.UUID); err != nil {
		global.LOGGER.Error("Fail to delete user!", zap.Error(err))
		response.FailWithMessage("Fail to delete user", c)
	} else {
		response.OkWithMessage("delete user success", c)
	}
}
