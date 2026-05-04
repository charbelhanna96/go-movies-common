package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/charbelhanna96/go-movies-common/pkg/middleware"
	"github.com/stretchr/testify/assert"
)

func dummyHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
}

func TestCORS_AllowedOrigin(t *testing.T) {
	handler := middleware.CORS([]string{"http://localhost:3000"}, dummyHandler())

	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Set("Origin", "http://localhost:3000")
	w := httptest.NewRecorder()

	handler.ServeHTTP(w, req)

	assert.Equal(t, "http://localhost:3000", w.Header().Get("Access-Control-Allow-Origin"))
	assert.Equal(t, "GET, OPTIONS", w.Header().Get("Access-Control-Allow-Methods"))
}

func TestCORS_DisallowedOrigin(t *testing.T) {
	handler := middleware.CORS([]string{"http://localhost:3000"}, dummyHandler())

	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Set("Origin", "http://evil.com")
	w := httptest.NewRecorder()

	handler.ServeHTTP(w, req)

	assert.Empty(t, w.Header().Get("Access-Control-Allow-Origin"))
}

func TestCORS_VaryHeaderAlwaysSet(t *testing.T) {
	handler := middleware.CORS([]string{"http://localhost:3000"}, dummyHandler())

	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Set("Origin", "http://evil.com")
	w := httptest.NewRecorder()

	handler.ServeHTTP(w, req)

	assert.Equal(t, "Origin", w.Header().Get("Vary"))
}

func TestCORS_OptionsReturns204(t *testing.T) {
	handler := middleware.CORS([]string{"http://localhost:3000"}, dummyHandler())

	req := httptest.NewRequest("OPTIONS", "/", nil)
	req.Header.Set("Origin", "http://localhost:3000")
	w := httptest.NewRecorder()

	handler.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
}

func TestCORS_NoOriginHeader(t *testing.T) {
	handler := middleware.CORS([]string{"http://localhost:3000"}, dummyHandler())

	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	handler.ServeHTTP(w, req)

	assert.Empty(t, w.Header().Get("Access-Control-Allow-Origin"))
	assert.Equal(t, "Origin", w.Header().Get("Vary"))
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCORS_MultipleAllowedOrigins(t *testing.T) {
	handler := middleware.CORS([]string{"http://localhost:3000", "http://localhost:8081"}, dummyHandler())

	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Set("Origin", "http://localhost:8081")
	w := httptest.NewRecorder()

	handler.ServeHTTP(w, req)

	assert.Equal(t, "http://localhost:8081", w.Header().Get("Access-Control-Allow-Origin"))
}
