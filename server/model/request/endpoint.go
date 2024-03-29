package request

type EndpointSearchQuery struct {
	Params EndpointSearchParam `json:"params" form:"params" binding:"required"`
}

type EndpointSearchParam struct {
	Pagination
	Desc        bool   `json:"desc" form:"desc"` // order by desc (by default)
	Name        string `json:"name" form:"name"`
	Method      string `json:"method" form:"method"`
	Path        string `json:"path" form:"path"`
	OrderBy     string `json:"orderBy" form:"orderBy" binding:"oneof=name path id method group_name"`
	Description string `json:"description" form:"description"`
	GroupName   string `json:"groupName" form:"groupName"`
}
