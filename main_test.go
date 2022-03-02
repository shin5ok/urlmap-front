package main

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var testg *gin.Engine

func TestPing(t *testing.T) {
	if testg == nil {
		testg = CreateRouter()
	}

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest(http.MethodGet, "/Ping", nil)

	testg.ServeHTTP(w, c.Request)

	assert.Equal(t, http.StatusOK, w.Code)
	// var obj struct {
	// 	Status  string
	// 	Version string
	// 	Message string
	// }
	// define struct and create object with it at the same time
	// c.BindJSON(&obj)
	// // assert.NoError(t, c.BindJSON(&obj))
	// log.Printf("%+v", obj)
	// assert.Equal(t, "Pong", obj.Message)
}

func TestRoot(t *testing.T) {
	if testg == nil {
		testg = CreateRouter()
	}

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest(http.MethodGet, "/", nil)

	testg.ServeHTTP(w, c.Request)

	assert.Equal(t, http.StatusOK, w.Code)
}

func Test_initDefaultResponse(t *testing.T) {
	ret := map[string]interface{}{"Message": "", "Status": "fail"}
	if got := initDefaultResponse(); !reflect.DeepEqual(got, ret) {
		t.Errorf("initDefaultResponse() = %v, want %v", got, ret)
	}
}
