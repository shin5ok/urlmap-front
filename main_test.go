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
	w := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(w)
	router := gin.Default()

	req, _ := http.NewRequest("GET", "/Ping", nil)
	context.Request = req
	router.GET("/", Ping)
	// router.ServeHTTP(w, req)
	t.Run("test /Ping", func(t *testing.T) {
		assert.Equal(t, http.StatusOK, w.Code)
	})
}
