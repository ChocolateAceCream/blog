package request

type PreviewArticle struct {
	Params FindById `json:"params" form:"params" binding:"required"`
}

type ArticleCursorListQuery struct {
	Params CursorListParam `json:"params" form:"params" binding:"required"`
}

type ArticleSearchQuery struct {
	Params ArticleSearchParam `json:"params" form:"params" binding:"required"`
}

type ArticleSearchParam struct {
	Pagination
	Keywords  string `json:"keywords" form:"keywords"`
	Published int    `json:"published" form:"published"`
	// Title    string `json:"title" form:"title"`
	// Author   string `json:"author" form:"author"`
	Desc bool `json:"desc" form:"desc"` // order by desc (by default)
}
