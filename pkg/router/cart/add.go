package cart

import (
    "net/http"
    "strconv"
    "time"

    "github.com/NTUT-Database-System-Course/ACW-Backend/pkg/config"
    "github.com/labstack/echo/v4"
)

// Handles adding a product to the user's cart
// @Summary Add a product to the user's cart
// @Description Add a product to the user's cart
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
// @Router /api/cart/add [post]
func Add(c echo.Context) error {
    userID := c.Get("user_id").(int)

	// Parse the product ID and count
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

	// Check if the product exists
	query := `SELECT id FROM product WHERE id = $1`
	var id int
	err = config.DB.QueryRow(query, productID).Scan(&id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Product does not exist",
		})
	}

	// Check if the product is already in the cart
	query = `SELECT 1 FROM cart WHERE member_id = $1 AND product_id = $2`
	var exists int = 0
	err = config.DB.QueryRow(query, userID, productID).Scan(&exists)
	if err == nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Product is already in cart",
		})
	}
	if exists != 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Product is already in cart",
		})
	}


	// Get current time in Taipei timezone
    location, err := time.LoadLocation("Asia/Taipei")
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{
            "error": "Failed to load location",
        })
    }
    now := time.Now().In(location)

	// Add the product to the cart
	query = `INSERT INTO cart (member_id, product_id, count, time) VALUES ($1, $2, $3, $4)`
	_, err = config.DB.Exec(query, userID, productID, count, now)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to add product to cart",
		})
	}

    return c.JSON(http.StatusOK, map[string]interface{}{
        "message": "Product added to cart",
    })
}