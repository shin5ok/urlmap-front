package main

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var testg = CreateRouter()

func TestPing(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/Ping", nil)

	testg.ServeHTTP(w, req)
	log.Println(http.StatusOK)
	log.Println(w.Code)
	assert.Equal(t, http.StatusOK, w.Code)
}
