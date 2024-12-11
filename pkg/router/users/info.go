package users

import (
    "fmt"
    "net/http"
    "strconv"

    "github.com/labstack/echo/v4"
    "github.com/NTUT-Database-System-Course/ACW-Backend/pkg/config"
)

// Info retrieves user information
// @Summary Get user info
// @Description Get a list of all users
// @Tags users
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Number of users per page" default(10)
// @Param sort query string false "Sort by field (id or name)" default(id)
// @Success 200 {array} UserResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/users/info [get]
func Info(c echo.Context) error {
    // Get query parameters
    pageParam := c.QueryParam("page")
    limitParam := c.QueryParam("limit")
    nameFilter := c.QueryParam("name")
    sortBy := c.QueryParam("sort")

    // Set default values
    page := 1
    limit := 10 // Default limit

    // Parse 'page' parameter
    if pageParam != "" {
        parsedPage, err := strconv.Atoi(pageParam)
        if err != nil || parsedPage < 1 {
            return c.JSON(http.StatusBadRequest, map[string]string{
                "error": "Invalid 'page' parameter; must be a positive integer",
            })
        }
        page = parsedPage
    }

    // Parse 'limit' parameter
    if limitParam != "" {
        parsedLimit, err := strconv.Atoi(limitParam)
        if err != nil || parsedLimit < 1 {
            return c.JSON(http.StatusBadRequest, map[string]string{
                "error": "Invalid 'limit' parameter; must be a positive integer",
            })
        }
        limit = parsedLimit
    }

    offset := (page - 1) * limit

    // Validate 'sort' parameter
    allowedSortFields := map[string]bool{
        "id":   true,
        "name": true,
    }
    if sortBy == "" {
        sortBy = "id" // Default sort field
    } else if !allowedSortFields[sortBy] {
        return c.JSON(http.StatusBadRequest, map[string]string{
            "error": "Invalid 'sort' parameter; allowed values are 'id' or 'name'",
        })
    }

    // Build the base query
    baseQuery := "SELECT id, name FROM users"
    var args []interface{}
    whereClauses := ""

    // Add name filtering if provided
    if nameFilter != "" {
        whereClauses = " WHERE name ILIKE $1"
        args = append(args, "%"+nameFilter+"%")
    }

    // Complete query with sorting and pagination
    query := fmt.Sprintf("%s%s ORDER BY %s LIMIT $%d OFFSET $%d", baseQuery, whereClauses, sortBy, len(args)+1, len(args)+2)
    args = append(args, limit, offset)

    // Fetch users
    rows, err := config.DB.Query(query, args...)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{
            "error": "Failed to query users",
        })
    }
    defer rows.Close()

    var users []UserResponse
    for rows.Next() {
        var user UserResponse
        if err := rows.Scan(&user.ID, &user.Name); err != nil {
            return c.JSON(http.StatusInternalServerError, map[string]string{
                "error": "Failed to scan user data",
            })
        }
        users = append(users, user)
    }

    // Check for errors after iterating
    if err = rows.Err(); err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{
            "error": "Error reading user data",
        })
    }

    // Return users
    return c.JSON(http.StatusOK, users)
}