package request

type EndpointSearchQuery struct {
	Params EndpointSearchParma `json:"params" form:"params" binding:"required"`
}

type EndpointSearchParma struct {
	Pagination
	Desc        bool   `json:"desc" form:"desc"` // order by desc (by default)
	Name        string `json:"name" form:"name"`
	Method      string `json:"method" form:"method"`
	Path        string `json:"path" form:"path"`
	OrderBy     string `json:"orderBy" form:"orderBy" binding:"oneof=name path id method"`
	Description string `json:"description" form:"description"`
	Group       string `json:"group" form:"group"`
}
