package test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
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

	// Test Member Info Success
	t.Run("Test Member Info Success", func(t *testing.T) {
		e := echo.New()
		router.NewRouter(e)

		// Simulate login to get token
		loginReq := httptest.NewRequest(http.MethodPost, "/api/auth/login", strings.NewReader(`{"username":"member","password":"password"}`))
		loginReq.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		loginRec := httptest.NewRecorder()
		e.ServeHTTP(loginRec, loginReq)

		assert.Equal(t, http.StatusOK, loginRec.Code)

		// Parse the login response to get the token
		var loginResponse struct {
			Message string `json:"message"`
			Token   string `json:"token"`
			Role    string `json:"role"`
		}
		err := json.Unmarshal(loginRec.Body.Bytes(), &loginResponse)
		assert.NoError(t, err)
		token := loginResponse.Token

		// Use token in Authorization header
		req := httptest.NewRequest(http.MethodGet, "/api/member/info", nil)
		req.Header.Set(echo.HeaderAuthorization, token)
		rec := httptest.NewRecorder()
		e.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)
	})
}
