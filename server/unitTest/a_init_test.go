/*
* @fileName init_test.go
* @author Di Sheng
* @date 2023/02/13 15:12:43
* @description test init api
	make sure this file name is before any other unit test files so this one run first
*/

package unitTest

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestAInitApi(t *testing.T) {
	r := RouterInstance
	req, _ := http.NewRequest("POST", "/api/public/initDB", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}
