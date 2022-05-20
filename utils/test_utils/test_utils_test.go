package test_utils

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetMockedContext(t *testing.T) {
	gin.SetMode(gin.TestMode)
	response := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/path", nil)

	c := GetMockedContext(request, response)

	assert.EqualValues(t, http.MethodGet, c.Request.Method)
	assert.EqualValues(t, "/path", c.Request.URL.Path)
}
