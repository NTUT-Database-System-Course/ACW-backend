package product

import (
    "log"
    "net/http"
    "time"
    "database/sql"

    "github.com/NTUT-Database-System-Course/ACW-Backend/pkg/config"
    "github.com/labstack/echo/v4"
)

// Create handles creating a new product
// @Summary Create product
// @Description Create a new product
// @Security ApiKeyAuth
// @Tags product
// @Accept json
// @Produce json
// @Param product body CreationRequest true "Create Product"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/product/create [post]
func Create(c echo.Context) error {
    userID := c.Get("user_id").(int)

    // Check if the user is a vendor
    query := `SELECT user_id FROM vendor WHERE user_id = $1`
    var vendorID int
    err := config.DB.QueryRow(query, userID).Scan(&vendorID)
    if err == sql.ErrNoRows {
        return c.JSON(http.StatusBadRequest, map[string]string{
            "error": "User is not a vendor",
        })
    } else if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{
            "error": "Failed to check user",
        })
    }

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
    buildTime := time.Now().In(location)

    // Insert product into the database
    query = `INSERT INTO product (price, description, name, remain, disability, image_url, build_time, vendor_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`
    var productID int
    err = tx.QueryRow(query, req.Price, req.Description, req.Name, req.Remain, req.Disability, req.ImageURL, buildTime, vendorID).Scan(&productID)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{
            "error": "Failed to create product",
        })
    }

    // Insert tags associated with the product
    for _, tagID := range req.Tags {
        tagQuery := `INSERT INTO own (product_id, tag_id) VALUES ($1, $2)`
        _, err := tx.Exec(tagQuery, productID, tagID)
        if err != nil {
            log.Printf("Failed to insert tag %d for product %d: %v", tagID, productID, err)
            return c.JSON(http.StatusInternalServerError, map[string]string{
                "error": "Failed to associate tags with product",
            })
        }
    }

    // Commit the transaction
    if err := tx.Commit(); err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{
            "error": "Failed to commit transaction",
        })
    }

    return c.JSON(http.StatusOK, map[string]interface{}{
        "message": "Product created successfully",
        "product_id": productID,
    })
}