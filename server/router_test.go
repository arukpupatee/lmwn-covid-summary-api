package server

import (
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
