package config_test

import (
	"os"
	"testing"

	"github.com/charbelhanna96/go-movies-common/pkg/config"
	"github.com/stretchr/testify/assert"
)

func TestGetEnv_ReturnsValue(t *testing.T) {
	os.Setenv("TEST_KEY", "hello")
	defer os.Unsetenv("TEST_KEY")
	assert.Equal(t, "hello", config.GetEnv("TEST_KEY", "fallback"))
}

func TestGetEnv_ReturnsFallbackWhenMissing(t *testing.T) {
	os.Unsetenv("TEST_KEY")
	assert.Equal(t, "fallback", config.GetEnv("TEST_KEY", "fallback"))
}

func TestGetEnv_ReturnsFallbackWhenEmpty(t *testing.T) {
	os.Setenv("TEST_KEY", "  ")
	defer os.Unsetenv("TEST_KEY")
	assert.Equal(t, "fallback", config.GetEnv("TEST_KEY", "fallback"))
}

func TestGetEnv_TrimsWhitespace(t *testing.T) {
	os.Setenv("TEST_KEY", "  hello  ")
	defer os.Unsetenv("TEST_KEY")
	assert.Equal(t, "hello", config.GetEnv("TEST_KEY", "fallback"))
}

func TestGetEnvInt_ReturnsValue(t *testing.T) {
	os.Setenv("TEST_INT", "42")
	defer os.Unsetenv("TEST_INT")
	assert.Equal(t, 42, config.GetEnvInt("TEST_INT", 0))
}

func TestGetEnvInt_ReturnsFallbackWhenMissing(t *testing.T) {
	os.Unsetenv("TEST_INT")
	assert.Equal(t, 10, config.GetEnvInt("TEST_INT", 10))
}

func TestGetEnvInt_ReturnsFallbackWhenInvalid(t *testing.T) {
	os.Setenv("TEST_INT", "abc")
	defer os.Unsetenv("TEST_INT")
	assert.Equal(t, 10, config.GetEnvInt("TEST_INT", 10))
}

func TestGetEnvInt_ReturnsFallbackWhenEmpty(t *testing.T) {
	os.Setenv("TEST_INT", "")
	defer os.Unsetenv("TEST_INT")
	assert.Equal(t, 10, config.GetEnvInt("TEST_INT", 10))
}

func TestGetEnvList_ReturnsValues(t *testing.T) {
	os.Setenv("TEST_LIST", "a,b,c")
	defer os.Unsetenv("TEST_LIST")
	assert.Equal(t, []string{"a", "b", "c"}, config.GetEnvList("TEST_LIST", nil))
}

func TestGetEnvList_ReturnsFallbackWhenMissing(t *testing.T) {
	os.Unsetenv("TEST_LIST")
	fallback := []string{"x", "y"}
	assert.Equal(t, fallback, config.GetEnvList("TEST_LIST", fallback))
}

func TestGetEnvList_TrimsWhitespace(t *testing.T) {
	os.Setenv("TEST_LIST", " a , b , c ")
	defer os.Unsetenv("TEST_LIST")
	assert.Equal(t, []string{"a", "b", "c"}, config.GetEnvList("TEST_LIST", nil))
}

func TestGetEnvList_SkipsEmptyParts(t *testing.T) {
	os.Setenv("TEST_LIST", "a,,b")
	defer os.Unsetenv("TEST_LIST")
	assert.Equal(t, []string{"a", "b"}, config.GetEnvList("TEST_LIST", nil))
}

func TestGetEnvList_ReturnsFallbackWhenAllEmpty(t *testing.T) {
	os.Setenv("TEST_LIST", ",,,")
	defer os.Unsetenv("TEST_LIST")
	fallback := []string{"x"}
	assert.Equal(t, fallback, config.GetEnvList("TEST_LIST", fallback))
}