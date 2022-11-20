package apiV1

type ApiGroup struct {
	UserApi UserApi
	AuthApi AuthApi
}

var ApiGroupInstance = new(ApiGroup)
