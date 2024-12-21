package member

import (
    "net/http"
    "log"
    "github.com/lib/pq"
    "github.com/NTUT-Database-System-Course/ACW-Backend/pkg/config"
    "github.com/labstack/echo/v4"
    "golang.org/x/crypto/bcrypt"
)

// Register registers a new member
// @Summary Register member
// @Description Register member
// @Tags member
// @Accept json
// @Produce json
// @Param member body RegistrationRequest true "Member Registration"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 409 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/member/register [post]
func Register(c echo.Context) error {
    var req RegistrationRequest

    // Bind the incoming JSON to the RegistrationRequest struct
    if err := c.Bind(&req); err != nil {
        log.Printf("Error binding member data: %v", err)
        return c.JSON(http.StatusBadRequest, map[string]string{
            "error": "Invalid request payload",
        })
    }

    // Create a Member instance from the request data
    member := Member{
        Name:     req.Name,
        Username: req.Username,
        Password: req.Password,
        Email:    req.Email,
    }

    // Hash the member's password using bcrypt
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(member.Password), bcrypt.DefaultCost)
    if err != nil {
        log.Printf("Error hashing password: %v", err)
        return c.JSON(http.StatusInternalServerError, map[string]string{
            "error": "Failed to process password",
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

    // Insert the member into the database
    // 1. create a user in the database
    query := `INSERT INTO "user" ("name", "username", "password") VALUES ($1, $2, $3) RETURNING "id"`
    var userID int
    err = tx.QueryRow(query, member.Name, member.Username, string(hashedPassword)).Scan(&userID)
    if err != nil {
        if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" { // 23505 is the PostgreSQL error code for unique violation
            log.Printf("Username already exists: %v", err)
            return c.JSON(http.StatusConflict, map[string]string{
                "error": "Username already exists",
            })
        }
        log.Printf("Error inserting member into database: %v", err)
        return c.JSON(http.StatusInternalServerError, map[string]string{
            "error": "Failed to insert member into database",
        })
    }

    // 2. create a member in the database
    query = `INSERT INTO "member" ("email", "user_id") VALUES ($1, $2)`
    _, err = tx.Exec(query, member.Email, userID)
    if err != nil {
        log.Printf("Error inserting member info into database: %v", err)
        return c.JSON(http.StatusInternalServerError, map[string]string{
            "error": "Failed to insert member info into database",
        })
    }

    // Commit the transaction
    if err := tx.Commit(); err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{
            "error": "Failed to commit transaction",
        })
    }

    // Respond with a success message and the new member's ID
    return c.JSON(http.StatusOK, map[string]interface{}{
        "message": "Member registered successfully",
        "id": userID,
    })
}