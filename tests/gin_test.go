package tests

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	gwm_app "github.com/slhmy/go-webmods/app"
	gwm_gin "github.com/slhmy/go-webmods/modules/gin"
)

func TestGin(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/error", gwm_gin.ErrorHandler, func(ginCtx *gin.Context) {
		_ = ginCtx.Error(errors.New("test error"))
		_ = ginCtx.Error(gwm_app.NewNotFoundError("test", "test").GinError())
	})

	req, err := http.NewRequest("GET", "/error", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Errorf("Expected status %d, got %d", http.StatusNotFound, rec.Code)
	}
	t.Log(rec.Body.String())
}
