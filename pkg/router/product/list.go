package product

import (
	"database/sql"
	"math/rand"
	"net/http"
	"time"

	"github.com/NTUT-Database-System-Course/ACW-Backend/pkg/config"
	"github.com/labstack/echo/v4"
)

// List handles fetching all products
// @Summary List products
// @Description Get all products
// @Tags product
// @Accept json
// @Produce json
// @Param id query int false "Product ID"
// @Param vendor_id query string false "Vendor ID"
// @Param random query bool false "Get random 5 items"
// @Success 200 {object} []Product
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/product/list [get]
func List(c echo.Context) error {
	productID := c.QueryParam("id")
	vendorID := c.QueryParam("vendor_id")
	random := c.QueryParam("random") == "true"

	var products []Product
	var query string
	var rows *sql.Rows
	var err error

	if productID != "" && vendorID != "" {
		query = `SELECT p.id, p.name, p.description, p.price, v.announcement, p.remain, p.disability, p.image_url, p.build_time 
                 FROM product p 
                 JOIN vendor v ON p.vendor_id = v.user_id 
                 WHERE p.id = $1 AND v.user_id = $2`
		rows, err = config.DB.Query(query, productID, vendorID)
	} else if productID != "" {
		query = `SELECT p.id, p.name, p.description, p.price, v.announcement, p.remain, p.disability, p.image_url, p.build_time 
                 FROM product p 
                 JOIN vendor v ON p.vendor_id = v.user_id 
                 WHERE p.id = $1`
		rows, err = config.DB.Query(query, productID)
	} else if vendorID != "" {
		query = `SELECT p.id, p.name, p.description, p.price, v.announcement, p.remain, p.disability, p.image_url, p.build_time 
                 FROM product p 
                 JOIN vendor v ON p.vendor_id = v.user_id 
                 WHERE v.user_id = $1`
		rows, err = config.DB.Query(query, vendorID)
	} else {
		query = `SELECT p.id, p.name, p.description, p.price, v.announcement, p.remain, p.disability, p.image_url, p.build_time 
                 FROM product p 
                 JOIN vendor v ON p.vendor_id = v.user_id`
		rows, err = config.DB.Query(query)
	}

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to fetch products",
		})
	}
	defer rows.Close()

	for rows.Next() {
		var product Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.VendorAnnouncement, &product.Remain, &product.Disability, &product.ImageURL, &product.BuildTime); err != nil {
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

	if random {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		r.Shuffle(len(products), func(i, j int) { products[i], products[j] = products[j], products[i] })
		if len(products) > 5 {
			products = products[:5]
		}
	}

	return c.JSON(http.StatusOK, products)
}
