package member

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/NTUT-Database-System-Course/ACW-Backend/pkg/config"
	"github.com/labstack/echo/v4"
)

// Update handles updating a member's info
// @Summary Update member info
// @Description Update member info
// @Security ApiKeyAuth
// @Tags member
// @Accept json
// @Produce json
// @Param update body UpdateRequest true "Update Request"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/member/update [put]
func Update(c echo.Context) error {
	userID := c.Get("user_id").(int)

	// Parse request
	var req UpdateRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request",
		})
	}

	// Start a transaction
	tx, err := config.DB.Begin()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to start transaction",
		})
	}
	defer tx.Rollback()

	// Check if the user is a member
	query := `SELECT user_id FROM member WHERE user_id = $1`
	var memberID int
	err = tx.QueryRow(query, userID).Scan(&memberID)
	if err == sql.ErrNoRows {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "User is not a member",
		})
	} else if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to check user",
		})
	}

	// Update member info
	query = `
		UPDATE "member"
		SET "email" = $1, "address" = $2, "phone_num" = $3
		WHERE "user_id" = $4
	`
	_, err = tx.Exec(query, req.Email, req.Address, req.PhoneNum, userID)
	if err != nil {
		log.Printf("Error updating member info: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to update member info",
		})
	}

	// Update user info
	query = `
		UPDATE "user"
		SET "name" = $1
		WHERE "id" = $2
	`
	_, err = tx.Exec(query, req.Name, userID)
	if err != nil {
		log.Printf("Error updating user info: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to update user info",
		})
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		log.Printf("Error committing transaction: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to commit transaction",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Member info updated successfully",
	})
}
