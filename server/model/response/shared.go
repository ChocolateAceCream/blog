/*
* @fileName shared.go
* @author Di Sheng
* @date 2022/10/18 08:49:42
* @description: wrap out http response
 */

package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int         `json:"errorCode"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

type Paging struct {
	List     interface{}
	Total    int64 `json:"total"`
	Page     int   `json:"page"`
	PageSize int   `json:"pageSize"`
}

const (
	ERROR        = 1
	SUCCESS      = 0
	UNAUTHORIZED = 401
)

func ResponseGenerator(code int, data interface{}, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}

func OkWithFullDetails(data interface{}, msg string, c *gin.Context) {
	ResponseGenerator(SUCCESS, data, msg, c)
}

func OkWithMessage(msg string, c *gin.Context) {
	ResponseGenerator(SUCCESS, map[string]interface{}{}, msg, c)
}

func FailWithMessage(msg string, c *gin.Context) {
	ResponseGenerator(ERROR, map[string]interface{}{}, msg, c)
}

func FailWithFullDetails(data interface{}, msg string, c *gin.Context) {
	ResponseGenerator(ERROR, data, msg, c)
}

func FailWithUnauthorized(msg string, c *gin.Context) {
	ResponseGenerator(UNAUTHORIZED, map[string]interface{}{}, msg, c)

}
