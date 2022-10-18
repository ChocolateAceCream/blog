package apiV1

type ApiGroup struct {
	UserApi UserApi
}

var ApiGroupInstance = new(ApiGroup)
