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
	userApi := apiV1.ApiGroupInstance.UserApi
	casbinApi := apiV1.ApiGroupInstance.CasbinApi
	authApi := apiV1.ApiGroupInstance.AuthApi
	roleApi := apiV1.ApiGroupInstance.RoleApi
	menuApi := apiV1.ApiGroupInstance.MenuApi
	initApi := apiV1.ApiGroupInstance.InitApi
	endpointApi := apiV1.ApiGroupInstance.EndpointApi
	articleApi := apiV1.ApiGroupInstance.ArticleApi
	ossApi := apiV1.ApiGroupInstance.OssApi
	commentApi := apiV1.ApiGroupInstance.CommentApi
	replyApi := apiV1.ApiGroupInstance.ReplyApi
	notificationApi := apiV1.ApiGroupInstance.NotificationApi
	PublicGroup := r.Group("/api/public")
	{
		// health check
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(200, "ok")
		})
		PublicGroup.POST("/initDB", initApi.InitDB)
	}

	auth := PublicGroup.Group("/auth")
	auth.Use(middleware.DefaultLimiter()).Use(middleware.SessionMiddleware())
	{
		auth.POST("/captcha", authApi.GetCaptcha)
		auth.POST("/sendEmailCode", authApi.SendEmailCode)
		auth.POST("/register", userApi.RegisterUser)
		auth.POST("/login", userApi.Login)
	}

	PrivateGroup := r.Group("")

	// turn on session
	PrivateGroup.Use(middleware.SessionMiddleware())

	// turn on sign verification
	if global.CONFIG.Signature.TurnOn {
		PrivateGroup.Use(middleware.SignVerifier())
	}

	// turn on casbin if gin mode is not test
	if gin.Mode() != "test" {
		PrivateGroup.Use(middleware.CasbinHandler())
	}

	v1 := PrivateGroup.Group("/api/v1")
	// v1.Use(middleware.Timer())
	{
		user := v1.Group("/user")
		// user.Use(middleware.DefaultLimiter())
		{
			user.GET("/list", userApi.GetUserList)
			user.POST("/active", userApi.ActiveUser)
			user.PUT("/resetPassword", userApi.ResetPassword)
			user.PUT("/edit", userApi.EditUser)
			user.DELETE("/delete", userApi.DeleteUser)
		}

		casbin := v1.Group("/casbin")
		{
			casbin.POST("/update", casbinApi.UpdateCasbin)
			casbin.GET("/list", casbinApi.GetCasbinByRoleId)
		}

		role := v1.Group("/role")
		{
			role.POST("/add", roleApi.AddRole)
			role.GET("/list", roleApi.GetRoleList)
			role.DELETE("/delete", roleApi.DeleteRole)
			role.PUT("/edit", roleApi.EditRole)

		}
		menu := v1.Group("/menu")
		{
			menu.POST("/add", menuApi.AddMenu)
			menu.GET("/currentUserMenu", menuApi.GetCurrentUserMenu)
			menu.POST("/getRoleMenuTree", menuApi.GetRoleMenuTree)
			menu.POST("/assignRoleMenus", menuApi.AssignRoleMenus)
			menu.GET("/list", menuApi.GetMenuList)
			menu.DELETE("/delete", menuApi.DeleteMenu)
			menu.PUT("/edit", menuApi.EditMenu)
		}

		endpoint := v1.Group("/endpoint")
		{
			endpoint.GET("/list", endpointApi.GetEndpointList)
			endpoint.POST("/add", endpointApi.AddEndpoint)
			endpoint.PUT("/edit", endpointApi.EditEndpoint)
			endpoint.DELETE("/delete", endpointApi.DeleteEndpoint)
		}

		article := v1.Group("/article")
		{
			article.GET("/list", articleApi.GetArticleList)
			article.GET("/search", articleApi.GetArticleSearchList)
			article.GET("/preview", articleApi.PreviewArticle)
			article.POST("/add", articleApi.AddArticle)
			article.PUT("/edit", articleApi.EditArticle)
			article.DELETE("/delete", articleApi.DeleteArticle)
		}

		comment := v1.Group("/comment")
		{
			comment.GET("/list", commentApi.GetCommentList)
			comment.POST("/add", commentApi.AddComment)
			comment.DELETE("/delete", commentApi.DeleteComment)
			comment.POST("/like", commentApi.LikeComment)
		}
		reply := v1.Group("/reply")
		{
			reply.GET("/list", replyApi.GetReplyList)
			reply.POST("/add", replyApi.AddReply)
			reply.DELETE("/delete", replyApi.DeleteReply)
			reply.POST("/like", replyApi.LikeReply)
		}

		notification := v1.Group("/notification")
		{
			notification.GET("/ws", notificationApi.WSHandler)
			notification.GET("/list", notificationApi.GetNotificationList)
			notification.GET("/unreadCount", notificationApi.GetUnreadCount)
			notification.DELETE("/delete", notificationApi.DeleteNotification)
			notification.PATCH("/read", notificationApi.ReadNotification)
		}

		oss := v1.Group("/oss")
		{
			oss.POST("/upload", ossApi.Uploader)
		}
	}
}
