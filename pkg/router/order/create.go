package order

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/NTUT-Database-System-Course/ACW-Backend/pkg/config"
	"github.com/labstack/echo/v4"
)

// Create handles creating a new order
// @Summary Create order
// @Description Create a new order
// @Security ApiKeyAuth
// @Tags order
// @Accept json
// @Produce json
// @Param order body CreationRequest true "Create Order"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/order/create [post]
func Create(c echo.Context) error {
	userID := c.Get("user_id").(int)

	// Check if the user is a member
	query := `SELECT user_id FROM member WHERE user_id = $1`
	var memberID int
	err := config.DB.QueryRow(query, userID).Scan(&memberID)
	if err == sql.ErrNoRows {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "User is not a member",
		})
	} else if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to check user",
		})
	}

	// Parse request
	var req CreationRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request",
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
			"error": "Failed to load timezone",
		})
	}
	currentTime := time.Now().In(location)

	// Insert order into the database
	query = `INSERT INTO "order" (name, description, state, address, phone_num, member_id, payment_method, shipment_method) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`
	var orderID int
	err = tx.QueryRow(query, req.Name, req.Description, req.State, req.Address, req.PhoneNum, memberID, req.PaymentMethod, req.ShipmentMethod).Scan(&orderID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to create order",
		})
	}

	// Get all products in the cart
	query = `SELECT product_id, count FROM cart WHERE member_id = $1`
	rows, err := config.DB.Query(query, memberID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to get products in cart",
		})
	}
	defer rows.Close()

	// Insert products into the list
	var isEmpty bool = true
	for rows.Next() {
		isEmpty = false
		var productID, count int
		if err := rows.Scan(&productID, &count); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "Failed to scan products",
			})
		}

		query = `INSERT INTO "list" ("order_id", "product_id", "count", "time") VALUES ($1, $2, $3, $4)`
		if _, err := tx.Exec(query, orderID, productID, count, currentTime); err != nil {
			log.Printf("orderID: %d, productID: %d, count: %d, time: %v", orderID, productID, count, currentTime)
			log.Printf("Failed to insert product %d into list: %v", productID, err)
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "Failed to insert product into list",
			})
		}

		// Update product stock and check if the stock is enough
		query = `SELECT remain FROM product WHERE id = $1`
		var stock int
		err = tx.QueryRow(query, productID).Scan(&stock)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "Failed to get product stock",
			})
		}
		if stock < count {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"error": "Not enough stock",
				"product_name": req.Name,
				"remain": stock,
			})
		}

		query = `UPDATE product SET remain = remain - $1 WHERE id = $2`
		if _, err := tx.Exec(query, count, productID); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "Failed to update product stock",
			})
		}
	}

	if isEmpty {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Cart is empty",
		})
	}

	// Clear the cart
	query = `DELETE FROM cart WHERE member_id = $1`
	if _, err := tx.Exec(query, memberID); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to clear cart",
		})
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to commit transaction",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Order created successfully",
	})
}
