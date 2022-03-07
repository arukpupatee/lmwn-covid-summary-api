package server

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func testSendHTTPRequest(req *http.Request) *httptest.ResponseRecorder {
	router := initRouter()
	res := httptest.NewRecorder()

	router.ServeHTTP(res, req)

	return res
}

func TestSuccessResponse(t *testing.T) {
	req, _ := http.NewRequest("GET", "/covid/summary", nil)
	res := testSendHTTPRequest(req)

	if res.Code != http.StatusOK {
		t.Fatalf("GET /covid/summary response %v, expect %v", res.Code, http.StatusOK)
	}
}

func TestResponseBodyFormat(t *testing.T) {
	req, _ := http.NewRequest("GET", "/covid/summary", nil)
	res := testSendHTTPRequest(req)

	resBody := Response{}

	err := json.Unmarshal(res.Body.Bytes(), &resBody)

	if err != nil {
		t.Fatalf("Error '%v' when parse response body from GET /covid/summary", err)
	}
	// TODO: Validate response body schema
}
