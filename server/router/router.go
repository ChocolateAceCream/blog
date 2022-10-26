/*
* @fileName router.go
* @author Di Sheng
* @date 2022/10/18 08:51:30
* @description init and load router
 */
package router

import (
	apiV1 "github.com/ChocolateAceCream/blog/api/v1"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RouterInit(r *gin.Engine) {
	// apply global middle ware
	// r.Use(middleware.GetMiddleware()...)

	//load router endpoints, basically wrapper over the code block below

	RouteLoader(r)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func RouteLoader(r *gin.Engine) {
	v1 := r.Group("/v1")
	{
		user := v1.Group("/user")
		userApi := apiV1.ApiGroupInstance.UserApi
		{
			user.GET("/userList", userApi.GetUserList)
			user.POST("/register", userApi.Register)
			user.PUT("/edit", userApi.EditUser)
			// user.PUT("/edit", EditUser)
			// user.DELETE("/delete", DeleteUser)
		}
	}

}
