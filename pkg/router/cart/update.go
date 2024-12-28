package cart

import (
    "net/http"
    "strconv"

    "github.com/NTUT-Database-System-Course/ACW-Backend/pkg/config"
    "github.com/labstack/echo/v4"
)

// Update handles updating the quantity of a product in the user's cart
// @Summary Update the quantity of a product in the user's cart
// @Description Update the quantity of a product in the user's cart
// @Security ApiKeyAuth
// @Tags cart
// @Accept json
// @Produce json
// @Param product_id query int true "Product ID"
// @Param count query int true "Count"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/cart/update [put]
func Update(c echo.Context) error {
    userID := c.Get("user_id").(int)

	// Parse the product ID and count from the query parameters
    productID, err := strconv.Atoi(c.QueryParam("product_id"))
    if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{
            "error": "Invalid product ID",
        })
    }
    count, err := strconv.Atoi(c.QueryParam("count"))
    if err != nil || count <= 0 {
        return c.JSON(http.StatusBadRequest, map[string]string{
            "error": "Invalid count",
        })
    }

	// Check if the product is in the user's cart
	query := `SELECT * FROM cart WHERE member_id = $1 AND product_id = $2`
	rows, err := config.DB.Query(query, userID, productID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to query cart",
		})
	}
	defer rows.Close()

	// If the product is not in the user's cart, return an error
	if !rows.Next() {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Product not in cart",
		})
	}

	// Update the quantity of the product in the user's cart
    query = `UPDATE cart SET count = $1 WHERE member_id = $2 AND product_id = $3`
    _, err = config.DB.Exec(query, count, userID, productID)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{
            "error": "Failed to update cart",
        })
    }

    return c.JSON(http.StatusOK, map[string]interface{}{
        "message": "Cart updated",
    })
}