package apiV1

type ApiGroup struct {
	UserApi     UserApi
	AuthApi     AuthApi
	CasbinApi   CasbinApi
	RoleApi     RoleApi
	MenuApi     MenuApi
	InitApi     InitApi
	EndpointApi EndpointApi
	ArticleApi  ArticleApi
	OssApi      OssApi
	CommentApi  CommentApi
}

var ApiGroupInstance = new(ApiGroup)
