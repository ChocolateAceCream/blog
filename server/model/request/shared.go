/*
* @fileName shared.go
* @author Di Sheng
* @date 2022/11/28 15:41:44
* @description shared request model
 */

package request

type Pagination struct {
	PageNumber int `json:"pageNumber" form:"pageNumber"`
	PageSize   int `json:"pageSize" form:"pageSize"`
}
