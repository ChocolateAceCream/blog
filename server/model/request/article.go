package request

type PreviewArticle struct {
	Params FindById `json:"params" form:"params" binding:"required"`
}

type ArticleCursorListQuery struct {
	Params ArticleCursorListParma `json:"params" form:"params" binding:"required"`
}

type ArticleCursorListParma struct {
	Pagination
	CursorId uint `json:"cursorId" form:"cursorId"`
	// Title    string `json:"title" form:"title"`
	// Author   string `json:"author" form:"author"`
	Desc bool `json:"desc" form:"desc"` // order by desc (by default)
}

type ArticleSearchQuery struct {
	Params ArticleSearchParma `json:"params" form:"params" binding:"required"`
}

type ArticleSearchParma struct {
	Pagination
	Keywords  string `json:"keywords" form:"keywords"`
	Published int    `json:"published" form:"published"`
	// Title    string `json:"title" form:"title"`
	// Author   string `json:"author" form:"author"`
	Desc bool `json:"desc" form:"desc"` // order by desc (by default)
}
