package favorite

import (
    "database/sql"
    "net/http"
    "strconv"
    "time"

    "github.com/NTUT-Database-System-Course/ACW-Backend/pkg/config"
    "github.com/labstack/echo/v4"
)

// Add handles adding a product to the user's favorite list
// @Summary Add favorite
// @Description Add a product to the user's favorite list
// @Security ApiKeyAuth
// @Tags favorite
// @Accept json
// @Produce json
// @Param product_id query int true "Product ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/favorite/add [post]
func Add(c echo.Context) error {
    userID := c.Get("user_id").(int)

    productID, err := strconv.Atoi(c.QueryParam("product_id"))
    if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{
            "error": "Invalid product ID",
        })
    }

    // Check if the product exists
    query := `SELECT id FROM product WHERE id = $1`
    var id int
    err = config.DB.QueryRow(query, productID).Scan(&id)
    if err == sql.ErrNoRows {
        return c.JSON(http.StatusBadRequest, map[string]string{
            "error": "Product does not exist",
        })
    } else if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{
            "error": "Failed to check product",
        })
    }

    // Check if the product is already in the favorite list
    query = `SELECT 1 FROM favor WHERE member_id = $1 AND product_id = $2`
    err = config.DB.QueryRow(query, userID, productID).Scan(new(int))
    if err == nil {
        return c.JSON(http.StatusBadRequest, map[string]string{
            "error": "Product is already in favorite",
        })
    } else if err != sql.ErrNoRows {
        return c.JSON(http.StatusInternalServerError, map[string]string{
            "error": "Failed to check favorite",
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

    // Get current time in Taipei timezone
    location, err := time.LoadLocation("Asia/Taipei")
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{
            "error": "Failed to load location",
        })
    }
    now := time.Now().In(location)

    // Add the product to the favorite list
    query = `INSERT INTO favor (member_id, product_id, time) VALUES ($1, $2, $3)`
    _, err = tx.Exec(query, userID, productID, now)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{
            "error": "Failed to add to favorite",
        })
    }

    // Commit the transaction
    err = tx.Commit()
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{
            "error": "Failed to commit transaction",
        })
    }

    return c.JSON(http.StatusOK, map[string]string{
        "message": "Product added to favorite",
    })
}