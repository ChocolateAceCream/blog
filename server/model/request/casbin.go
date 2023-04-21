package request

type CasbinPolicy struct {
	Path   string `json:"path"`
	Method string `json:"method"`
}

type GetCasbinByRole struct {
	Params FindById `json:"params" form:"params" binding:"required"`
}

type UpdateCasbin struct {
	RoleId       uint           `json:"roleId" binding:"required"`
	CasbinPolicy []CasbinPolicy `json:"endpoints"`
}

func DefaultCasbin() []CasbinPolicy {
	return []CasbinPolicy{
		//TODO: add more default routes
		{Path: "/api/v1/user/list", Method: "GET"},
		{Path: "/api/v1/user/resetPassword", Method: "PUT"},
		{Path: "/api/v1/user/edit", Method: "PUT"},
		{Path: "/api/v1/user/active", Method: "POST"},
		{Path: "/api/public/auth/register", Method: "POST"},
	}
}
