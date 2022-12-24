package request

type Endpoint struct {
	Path   string `json:"path"`
	Method string `json:"method"`
}

type UpdateCasbin struct {
	RoleId    uint       `json:"roleId" binding:"required"`
	Endpoints []Endpoint `json:"endpoints"`
}

func DefaultCasbin() []Endpoint {
	return []Endpoint{
		//TODO: add more default routes
		{Path: "/api/v1/user/userList", Method: "GET"},
	}
}
