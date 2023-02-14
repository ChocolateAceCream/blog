package request

type EndpointSearchQuery struct {
	Pagination
	Desc    bool   `json:"desc"` // order by desc (by default)
	PID     uint   `json:"pid"`
	Name    string `json:"name"`
	Type    int    `json:"type"`
	Method  string `json:"method"`
	Path    string `json:"path"`
	OrderBy string `json:"orderBy" form:"orderBy" binding:"oneof=name pid path id"`
}
