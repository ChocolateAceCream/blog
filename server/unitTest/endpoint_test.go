package unitTest

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ChocolateAceCream/blog/model/dbTable"
	"github.com/go-playground/assert/v2"
)

const GROUP = 1
const ENDPOINT = 2

func TestNewEndpointApis(t *testing.T) {
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
	t.Logf("----TestNewEndpointApi done----")
}
