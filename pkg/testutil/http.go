package testutil

import (
	"net/http"
	"net/http/httptest"
	"net/url"
)

// NewRequest creates a test HTTP request with optional query parameters.
func NewRequest(method, path string, params map[string]string) *http.Request {
	query := url.Values{}
	for k, v := range params {
		query.Set(k, v)
	}
	target := path
	if len(params) > 0 {
		target = path + "?" + query.Encode()
	}
	req := httptest.NewRequest(method, target, nil)
	return req
}

// NewRecorder creates a new response recorder.
func NewRecorder() *httptest.ResponseRecorder {
	return httptest.NewRecorder()
}
