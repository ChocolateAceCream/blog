package request

type EndpointSearchQuery struct {
	Pagination
	Desc        bool   `json:"desc"` // order by desc (by default)
	Name        string `json:"name"`
	Method      string `json:"method"`
	Path        string `json:"path"`
	OrderBy     string `json:"orderBy" form:"orderBy" binding:"oneof=name path id method"`
	Description string `json:"description"`
	Group       string `json:"group"`
}
