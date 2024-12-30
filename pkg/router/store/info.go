package store

import (
	"database/sql"
	"net/http"

	"github.com/NTUT-Database-System-Course/ACW-Backend/pkg/config"
	"github.com/labstack/echo/v4"
)

// Info gets the store info
// @Summary Get store info
// @Description Get store info
// @Security ApiKeyAuth
// @Tags store
// @Accept json
// @Produce json
// @Success 200 {object} StoreInfo
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/store/info [get]
func Info(c echo.Context) error {
	userID := c.Get("user_id").(int)

	// Check if the user is a store
	var storeID int
	query := `SELECT "user_id" FROM "vendor" WHERE "user_id" = $1`
	err := config.DB.QueryRow(query, userID).Scan(&storeID)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"error": "User is not a store",
			})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to fetch store",
		})
	}

	// Parse the store info
	storeInfo := StoreInfo{
		ID: storeID,
	}

	return c.JSON(http.StatusOK, storeInfo)
}
