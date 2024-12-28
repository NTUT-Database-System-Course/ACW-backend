package cart

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/NTUT-Database-System-Course/ACW-Backend/pkg/config"
	"github.com/labstack/echo/v4"
)

// Delete handles removing a product from the user's cart
// @Summary Remove a product from the user's cart
// @Description Remove a product from the user's cart
// @Security ApiKeyAuth
// @Tags cart
// @Accept json
// @Produce json
// @Param product_id query int true "Product ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/cart/delete [delete]
func Delete(c echo.Context) error {
	userID := c.Get("user_id").(int)

	// Parse the product ID from the query parameters
	productID, err := strconv.Atoi(c.QueryParam("product_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid product ID",
		})
	}

	// Check if the product is in the user's cart
	query := `SELECT 1 FROM cart WHERE member_id = $1 AND product_id = $2`
	var exists int = 0
	err = config.DB.QueryRow(query, userID, productID).Scan(&exists)
	if err != sql.ErrNoRows {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Product not in cart",
		})
	}

	// Remove the product from the user's cart
	query = `DELETE FROM cart WHERE member_id = $1 AND product_id = $2`
	_, err = config.DB.Exec(query, userID, productID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to remove from cart",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Product removed from cart",
	})
}
