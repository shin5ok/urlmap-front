package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	// "github.com/go-playground/assert/v2"
	"github.com/stretchr/testify/assert"
)

func TestPing(t *testing.T) {
	recorder := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(recorder)

	req, _ := http.NewRequest("GET", "/Ping", nil)
	context.Request = req
	Ping(context)
	// asserts := assert.New(t)
	t.Run("test /Ping", func(t *testing.T) {
		assert.Equal(t, http.StatusOK, recorder.Code)
	})
}
