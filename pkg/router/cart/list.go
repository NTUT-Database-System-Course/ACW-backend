package cart

import (
    "net/http"

    "github.com/NTUT-Database-System-Course/ACW-Backend/pkg/config"
    "github.com/labstack/echo/v4"
)

// List handles listing all products in the user's cart
// @Summary List cart items
// @Description List all products in the user's cart
// @Security ApiKeyAuth
// @Tags cart
// @Accept json
// @Produce json
// @Success 200 {array} CartItem
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/cart/list [get]
func List(c echo.Context) error {
    userID := c.Get("user_id").(int)

	// Get all products in the user's cart
    query := `
		SELECT p.id, p.name, p.description, p.price, p.vendor_id, p.remain, p.disability, p.build_time, p.image_url, c.count, c.time
		FROM cart c
		JOIN product p ON c.product_id = p.id
		WHERE c.member_id = $1
	`
	rows, err := config.DB.Query(query, userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to fetch cart items",
		})
	}

	// Parse the rows
	var cartItems []CartItem
	for rows.Next() {
		var cartItem CartItem
		err = rows.Scan(&cartItem.ID, &cartItem.Name, &cartItem.Description, &cartItem.Price, &cartItem.VendorID, &cartItem.Remain, &cartItem.Disability, &cartItem.BuildTime, &cartItem.ImageURL, &cartItem.Count, &cartItem.CreationTime)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "Failed to parse cart items",
			})
		}
		// get tags
		query = `SELECT tag_id FROM own WHERE product_id = $1`
		rows, err := config.DB.Query(query, cartItem.ID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "Failed to fetch tags",
			})
		}
		for rows.Next() {
			var tagID int
			err = rows.Scan(&tagID)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{
					"error": "Failed to parse tags",
				})
			}
			cartItem.Tags = append(cartItem.Tags, tagID)
		}
		cartItems = append(cartItems, cartItem)
	}

	// Return the cart items
	return c.JSON(http.StatusOK, cartItems)
}