package request

type PreviewArticle struct {
	Params FindById `json:"params" form:"params" binding:"required"`
}
