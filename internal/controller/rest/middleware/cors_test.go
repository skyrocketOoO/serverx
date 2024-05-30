package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"web-server-template/internal/controller/rest/middleware"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCORS(t *testing.T) {
	r := gin.New()
	r.Use(middleware.CORS())

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "")
	})

	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, "*", w.Header().Get("Access-Control-Allow-Origin"))
}
