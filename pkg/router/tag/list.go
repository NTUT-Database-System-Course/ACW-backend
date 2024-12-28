package tag

import (
	"github.com/NTUT-Database-System-Course/ACW-Backend/pkg/config"
	"github.com/labstack/echo/v4"
	"net/http"
)

// List all tags
// @Summary List all tags
// @Description List all tags
// @Tags tag
// @Accept json
// @Produce json
// @Success 200 {array} Tag
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/tag/list [get]
func List(c echo.Context) error {
	var tags []Tag
	query := `SELECT "id", "name", "type" FROM "tag"`
	rows, err := config.DB.Query(query)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to fetch tags",
		})
	}
	defer rows.Close()

	for rows.Next() {
		var tag Tag
		if err := rows.Scan(&tag.ID, &tag.Name, &tag.Type); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "Failed to scan tag",
			})
		}
		tags = append(tags, tag)
	}

	return c.JSON(http.StatusOK, tags)
}
