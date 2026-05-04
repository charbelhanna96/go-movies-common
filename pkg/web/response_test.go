package web_test

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/charbelhanna96/go-movies-common/pkg/testutil"
	"github.com/charbelhanna96/go-movies-common/pkg/web"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestJSON_WritesStatusAndBody(t *testing.T) {
	w := testutil.NewRecorder()
	web.JSON(w, http.StatusOK, map[string]string{"key": "value"})

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "application/json", w.Header().Get("Content-Type"))

	var result map[string]string
	require.NoError(t, json.NewDecoder(w.Body).Decode(&result))
	assert.Equal(t, "value", result["key"])
}

func TestJSON_Writes201(t *testing.T) {
	w := testutil.NewRecorder()
	web.JSON(w, http.StatusCreated, map[string]string{"status": "created"})
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestJSON_WritesEmptyArray(t *testing.T) {
	w := testutil.NewRecorder()
	web.JSON(w, http.StatusOK, []string{})

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "[]")
}

func TestError_WritesMessageAndStatus(t *testing.T) {
	w := testutil.NewRecorder()
	web.Error(w, http.StatusBadRequest, "invalid input")

	assert.Equal(t, http.StatusBadRequest, w.Code)

	var result map[string]string
	require.NoError(t, json.NewDecoder(w.Body).Decode(&result))
	assert.Equal(t, "invalid input", result["message"])
}

func TestError_Writes500(t *testing.T) {
	w := testutil.NewRecorder()
	web.Error(w, http.StatusInternalServerError, "internal server error")

	assert.Equal(t, http.StatusInternalServerError, w.Code)

	var result map[string]string
	require.NoError(t, json.NewDecoder(w.Body).Decode(&result))
	assert.Equal(t, "internal server error", result["message"])
}
