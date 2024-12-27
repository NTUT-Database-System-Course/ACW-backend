package favorite

import (
    "database/sql"
    "net/http"
    "strconv"

    "github.com/NTUT-Database-System-Course/ACW-Backend/pkg/config"
    "github.com/labstack/echo/v4"
)

// Delete handles removing a product from the user's favorite list
// @Summary Delete favorite
// @Description Remove a product from the user's favorite list
// @Security ApiKeyAuth
// @Tags favorite
// @Accept json
// @Produce json
// @Param product_id query int true "Product ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/favorite/delete [delete]
func Delete(c echo.Context) error {
    userID := c.Get("user_id").(int)

    productID, err := strconv.Atoi(c.QueryParam("product_id"))
    if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{
            "error": "Invalid product ID",
        })
    }

    // Check if the product is in the favorite list
    query := `SELECT 1 FROM favor WHERE member_id = $1 AND product_id = $2`
    err = config.DB.QueryRow(query, userID, productID).Scan(new(int))
    if err == sql.ErrNoRows {
        return c.JSON(http.StatusBadRequest, map[string]string{
            "error": "Product is not in favorite",
        })
    } else if err != nil {
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

    // Remove the product from the favorite list
    query = `DELETE FROM favor WHERE member_id = $1 AND product_id = $2`
    _, err = tx.Exec(query, userID, productID)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{
            "error": "Failed to remove from favorite",
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
        "message": "Product removed from favorite",
    })
}