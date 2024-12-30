package order

import (
	"database/sql"
	"net/http"

	"github.com/NTUT-Database-System-Course/ACW-Backend/pkg/config"
	"github.com/labstack/echo/v4"
)

// List handles listing all orders
// @Summary List orders
// @Description List all orders
// @Security ApiKeyAuth
// @Tags order
// @Accept json
// @Produce json
// @Success 200 {object} []Order
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/order/list [get]
func List(c echo.Context) error {
	userID := c.Get("user_id").(int)

	// Check if the user is a member
	query := `SELECT user_id FROM member WHERE user_id = $1`
	var memberID int
	err := config.DB.QueryRow(query, userID).Scan(&memberID)
	if err == sql.ErrNoRows {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "User is not a member",
		})
	}

	// Get all orders
	query = `SELECT id, name, description, state, address, phone_num, payment_method, shipment_method FROM "order" WHERE member_id = $1`
	rows, err := config.DB.Query(query, memberID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to query orders",
		})
	}
	defer rows.Close()

	// Fetch all orders
	orders := []Order{}
	for rows.Next() {
		var order Order
		err := rows.Scan(&order.ID, &order.Name, &order.Description, &order.State, &order.Address, &order.PhoneNum, &order.PaymentMethod, &order.ShipmentMethod)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "Failed to scan orders",
			})
		}

		// Get all products in the order
		query = `
			SELECT p.name, p.description, p.price, u.name, p.build_time, p.image_url, l.count
			FROM list l
			JOIN product p ON l.product_id = p.id
			JOIN vendor v ON p.vendor_id = v.user_id
			JOIN "user" u ON v.user_id = u.id
			WHERE l.order_id = $1
		`
		rows, err := config.DB.Query(query, order.ID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "Failed to get products in order",
			})
		}
		defer rows.Close()

		// Fetch all products
		products := []Product{}
		for rows.Next() {
			var product Product
			err := rows.Scan(&product.Name, &product.Description, &product.Price, &product.VendorName, &product.BuildTime, &product.ImageURL, &product.Count)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{
					"error": "Failed to scan products",
				})
			}
			products = append(products, product)
		}
		order.Products = products
		orders = append(orders, order)
	}

	return c.JSON(http.StatusOK, orders)
}
