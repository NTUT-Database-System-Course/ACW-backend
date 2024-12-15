package users

import (
    "net/http"

    "github.com/labstack/echo/v4"
    "github.com/NTUT-Database-System-Course/ACW-Backend/pkg/config"
    "github.com/lib/pq"
    "golang.org/x/crypto/bcrypt"
    "log"
)

// Register registers a new user
// @Summary Register user
// @Description Register user
// @Tags users
// @Accept json
// @Produce json
// @Param user body UserRegistrationRequest true "User Registration"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/users/register [post]
func Register(c echo.Context) error {
    var req UserRegistrationRequest

    // Bind the incoming JSON to the UserRegistrationRequest struct
    if err := c.Bind(&req); err != nil {
        log.Printf("Error binding user data: %v", err)
        return c.JSON(http.StatusBadRequest, map[string]string{
            "error": "Invalid request payload",
        })
    }

    // Create a User instance from the request data
    user := User{
        Name:     req.Name,
        Username: req.Username,
        Password: req.Password,
    }

    // Hash the user's password using bcrypt
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        log.Printf("Error hashing password: %v", err)
        return c.JSON(http.StatusInternalServerError, map[string]string{
            "error": "Failed to process password",
        })
    }

    // Insert the user into the database
    query := `INSERT INTO users (name, username, password) VALUES ($1, $2, $3) RETURNING id`
    var userID int
    err = config.DB.QueryRow(query, user.Name, user.Username, string(hashedPassword)).Scan(&userID)
    if err != nil {
        log.Printf("Error inserting user into database: %v", err)
        if pgErr, ok := err.(*pq.Error); ok && pgErr.Code == "23505" {
            return c.JSON(http.StatusBadRequest, map[string]string{
                "error": "Username already exists",
            })
        }
        return c.JSON(http.StatusInternalServerError, map[string]string{
            "error": "Failed to register user",
        })
    }

    // Respond with a success message and the new user's ID
    return c.JSON(http.StatusOK, map[string]interface{}{
        "message": "User registered successfully",
        "user_id": userID,
    })
}