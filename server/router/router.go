/*
* @fileName router.go
* @author Di Sheng
* @date 2022/10/18 08:51:30
* @description init and load router
 */
package router

import (
	apiV1 "github.com/ChocolateAceCream/blog/api/v1"
	"github.com/ChocolateAceCream/blog/global"
	"github.com/ChocolateAceCream/blog/middleware"
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
	PublicGroup := r.Group("/api/public")
	{
		// health check
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(200, "ok")
		})
	}
	v1 := r.Group("/api/v1")
	v1.Use(middleware.SessionMiddleware())

	// turn on sign verification
	if global.CONFIG.Signature.TurnOn {
		v1.Use(middleware.SignVerifier())
	}
	// v1.Use(middleware.Timer()).Use(middleware.SessionMiddleware())
	{
		user := v1.Group("/user")
		// user.Use(middleware.DefaultLimiter())
		userApi := apiV1.ApiGroupInstance.UserApi
		{
			user.GET("/userList", userApi.GetUserList)
			user.POST("/register", userApi.RegisterUser)
			user.POST("/active", userApi.ActiveUser)
			user.PUT("/edit", userApi.EditUser)
			user.DELETE("/delete", userApi.DeleteUser)
		}

		auth := v1.Group("/auth")
		auth.Use(middleware.DefaultLimiter())
		authApi := apiV1.ApiGroupInstance.AuthApi
		{
			auth.POST("/captcha", authApi.GetCaptcha)
			auth.POST("/sendEmailCode", authApi.SendEmailCode)
		}
	}

}
