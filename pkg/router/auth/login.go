package auth

import (
	"log"
	"net/http"

	"github.com/NTUT-Database-System-Course/ACW-Backend/pkg/config"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

// Login logs in a user
// @Summary Login user
// @Description Login a user
// @Tags auth
// @Accept json
// @Produce json
// @Param login body LoginRequest true "Login Request"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/auth/login [post]
func Login(c echo.Context) error {
	var req LoginRequest

	// Bind the incoming JSON to the LoginRequest struct
	if err := c.Bind(&req); err != nil {
		log.Printf("Error binding login data: %v", err)
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request payload",
		})
	}

	// Fetch the user from the database
	var id int
	var password string
	query := `SELECT "id", "password" FROM "user" WHERE "username" = $1`
	err := config.DB.QueryRow(query, req.Username).Scan(&id, &password)
	if err != nil {
		log.Printf("Error fetching user from database: %v", err)
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"error": "Invalid username or password",
		})
	}

	// Compare the provided password with the stored hashed password
	if err := bcrypt.CompareHashAndPassword([]byte(password), []byte(req.Password)); err != nil {
		log.Printf("Invalid password: %v", err)
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"error": "Invalid username or password",
		})
	}

	// Generate JWT token
	token, err := config.GenerateJWT(id)
	if err != nil {
		log.Printf("Error generating JWT token: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to generate token",
		})
	}

	// Respond with the JWT token
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Login successful",
		"token":   token,
	})
}
