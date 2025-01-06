package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/NTUT-Database-System-Course/ACW-Backend/pkg/config"
	"github.com/NTUT-Database-System-Course/ACW-Backend/pkg/router"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestMemberInfo(t *testing.T) {
	// Initialize the database (adjust as needed for test DB/mocks)
	config.InitDB()

	// Test Member Info Unauthorized
	t.Run("Test Member Info Unauthorized", func(t *testing.T) {
		e := echo.New()
		router.NewRouter(e)

		// without Authorization header
		req := httptest.NewRequest(http.MethodGet, "/api/member/info", nil)
		rec := httptest.NewRecorder()
		e.ServeHTTP(rec, req)
		assert.Equal(t, http.StatusUnauthorized, rec.Code)
	})

}
