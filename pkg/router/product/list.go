package product

import (
    "net/http"
    "github.com/labstack/echo/v4"
    "github.com/NTUT-Database-System-Course/ACW-Backend/pkg/config"
)

// List handles fetching all products
// @Summary List products
// @Description Get all products
// @Tags product
// @Accept json
// @Produce json
// @Success 200 {object} []Product
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/product/list [get]
func List(c echo.Context) error {
    var products []Product
    query := `SELECT id, name, description, price, vendor_id, remain, disability, image_url, build_time FROM product`
    rows, err := config.DB.Query(query)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{
            "error": "Failed to fetch products",
        })
    }
    defer rows.Close()

    for rows.Next() {
        var product Product
        if err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.VendorID, &product.Remain, &product.Disability, &product.ImageURL, &product.BuildTime); err != nil {
            return c.JSON(http.StatusInternalServerError, map[string]string{
                "error": "Failed to scan product",
            })
        }

        // Fetch tags for the product
        tagQuery := `SELECT "tag_id" FROM "own" WHERE "product_id" = $1`
        tagRows, err := config.DB.Query(tagQuery, product.ID)
        if err != nil {
            return c.JSON(http.StatusInternalServerError, map[string]string{
                "error": "Failed to fetch product tags",
            })
        }
        defer tagRows.Close()

        var tags []int
        for tagRows.Next() {
            var tagID int
            if err := tagRows.Scan(&tagID); err != nil {
                return c.JSON(http.StatusInternalServerError, map[string]string{
                    "error": "Failed to scan tag",
                })
            }
            tags = append(tags, tagID)
        }
        product.Tags = tags

        products = append(products, product)
    }

    return c.JSON(http.StatusOK, products)
}