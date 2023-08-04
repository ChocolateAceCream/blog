/*
* @fileName shared.go
* @author Di Sheng
* @date 2022/11/28 15:41:44
* @description shared request model
 */

package request

type CursorListParam struct {
	Pagination
	CursorId uint `json:"cursorId" form:"cursorId"`
	// Title    string `json:"title" form:"title"`
	// Author   string `json:"author" form:"author"`
	Desc bool `json:"desc" form:"desc"` // order by desc (by default)
}

type Pagination struct {
	PageNumber int `json:"pageNumber" form:"pageNumber"`
	PageSize   int `json:"pageSize" form:"pageSize"`
}

type FindById struct {
	ID int `json:"id" form:"id"`
}

type FindByIds struct {
	ID []int `json:"id" form:"id"`
}
