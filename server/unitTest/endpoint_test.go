package unitTest

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ChocolateAceCream/blog/model/dbTable"
	"github.com/ChocolateAceCream/blog/model/request"
	"github.com/ChocolateAceCream/blog/model/response"
	"github.com/go-playground/assert/v2"
)

const GROUP = 1
const ENDPOINT = 2

func TestNewEndpointApi(t *testing.T) {
	r := RouterInstance
	payload := dbTable.Endpoint{
		Description: "Test Menu",
		Name:        "Test Menu",
		Method:      "POST",
		Path:        "/api/v1/test",
		PID:         3,
		Type:        ENDPOINT,
	}
	body, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", "/api/v1/endpoint/new", bytes.NewReader(body))
	req.Host = "localhost:1234"
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	// t.Logf("----TestNewEndpointApi done----")
}

func TestGetEndpointListApi(t *testing.T) {
	r := RouterInstance
	payload := request.EndpointSearchQuery{
		Pagination: request.Pagination{PageNumber: 1, PageSize: 1},
		Type:       1,
		OrderBy:    "name",
		Desc:       true,
	}
	body, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", "/api/v1/endpoint/list", bytes.NewReader(body))
	req.Host = "localhost:1234"
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	result, _ := ioutil.ReadAll(w.Body)
	m := &response.Response{}
	err := json.Unmarshal(result, &m)
	assert.Equal(t, err, nil)

	data, _ := json.Marshal(m.Data.(map[string]interface{})["List"])
	records := []dbTable.Endpoint{}
	_ = json.Unmarshal(data, &records)

	assert.Equal(t, records[0].Name, "User Group")
}
