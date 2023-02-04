package apiV1

type ApiGroup struct {
	UserApi   UserApi
	AuthApi   AuthApi
	CasbinApi CasbinApi
	RoleApi   RoleApi
	MenuApi   MenuApi
	InitApi   InitApi
}

var ApiGroupInstance = new(ApiGroup)
