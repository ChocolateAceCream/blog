package apiV1

type ApiGroup struct {
	UserApi   UserApi
	AuthApi   AuthApi
	CasbinApi CasbinApi
	RoleApi   RoleApi
}

var ApiGroupInstance = new(ApiGroup)
