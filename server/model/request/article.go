package request

type PreviewArticle struct {
	Params FindById `json:"params" form:"params" binding:"required"`
}

type ArticleSearchQuery struct {
	Params ArticleSearchParma `json:"params" form:"params" binding:"required"`
}

type ArticleSearchParma struct {
	Pagination
	CursorId uint `json:"cursorId" form:"cursorId"`
	// Title    string `json:"title" form:"title"`
	// Author   string `json:"author" form:"author"`
	Desc bool `json:"desc" form:"desc"` // order by desc (by default)
}
