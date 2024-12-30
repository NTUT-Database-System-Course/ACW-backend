package favorite

import (
	"net/http"

	"github.com/NTUT-Database-System-Course/ACW-Backend/pkg/config"
	"github.com/labstack/echo/v4"
)

// List handles listing the user's favorite products
// @Summary List favorites
// @Description List the user's favorite products
// @Security ApiKeyAuth
// @Tags favorite
// @Accept json
// @Produce json
// @Success 200 {array} FavoriteProduct
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/favorite/list [get]
func List(c echo.Context) error {
	userID := c.Get("user_id").(int)

	query := `SELECT p.id, p.name, p.description, p.price, p.remain, p.disability, p.image_url 
              FROM product p 
              JOIN favor f ON p.id = f.product_id 
              WHERE f.member_id = $1`
	rows, err := config.DB.Query(query, userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to list favorites",
		})
	}
	defer rows.Close()

	var favorites []FavoriteProduct
	for rows.Next() {
		var favorite FavoriteProduct
		err := rows.Scan(&favorite.ID, &favorite.Name, &favorite.Description, &favorite.Price, &favorite.Remain, &favorite.Disability, &favorite.ImageURL)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "Failed to scan favorite",
			})
		}
		favorites = append(favorites, favorite)
	}

	return c.JSON(http.StatusOK, favorites)
}
