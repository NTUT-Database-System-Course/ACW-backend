package product

import (
    "log"
    "net/http"
    "database/sql"

    "github.com/NTUT-Database-System-Course/ACW-Backend/pkg/config"
    "github.com/labstack/echo/v4"
)

// Update handles updating an existing product
// @Summary Update product
// @Description Update an existing product
// @Security ApiKeyAuth
// @Tags product
// @Accept json
// @Produce json
// @Param product body UpdateRequest true "Update Product"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/product/update [put]
func Update(c echo.Context) error {
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

    var req UpdateRequest
    if err := c.Bind(&req); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{
            "error": "Invalid request",
        })
    }
    
    // Check if the product exists and is owned by the vendor
    var productID int
    query = `SELECT id FROM product WHERE id = $1 AND vendor_id = $2`
    err = config.DB.QueryRow(query, req.ID, vendorID).Scan(&productID)
    if err == sql.ErrNoRows {
        return c.JSON(http.StatusBadRequest, map[string]string{
            "error": "Product does not exist or is not owned by the vendor",
        })
    } else if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{
            "error": "Failed to check product",
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

    // Update product in the database
    query = `UPDATE product SET price = $1, description = $2, name = $3, remain = $4, disability = $5, image_url = $6 WHERE id = $7`
    _, err = tx.Exec(query, req.Price, req.Description, req.Name, req.Remain, req.Disability, req.ImageURL, req.ID)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{
            "error": "Failed to update product",
        })
    }

    // Delete existing tags associated with the product
    deleteTagQuery := `DELETE FROM own WHERE product_id = $1`
    _, err = tx.Exec(deleteTagQuery, req.ID)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{
            "error": "Failed to delete existing tags",
        })
    }

    // Insert new tags associated with the product
    for _, tagID := range req.Tags {
        tagQuery := `INSERT INTO own (product_id, tag_id) VALUES ($1, $2)`
        _, err := tx.Exec(tagQuery, req.ID, tagID)
        if err != nil {
            log.Printf("Failed to insert tag %d for product %d: %v", tagID, req.ID, err)
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
        "message": "Product updated successfully",
    })
}