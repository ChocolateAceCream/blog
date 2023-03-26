package request

type FindById struct {
	ID int `json:"id" form:"id"`
}

type FindByIds struct {
	ID []int `json:"id" form:"id"`
}
