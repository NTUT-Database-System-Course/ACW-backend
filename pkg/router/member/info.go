package member

import (
	"github.com/NTUT-Database-System-Course/ACW-Backend/pkg/config"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

// Info gets the member info
// @Summary Get member info
// @Description Get member info
// @Security ApiKeyAuth
// @Tags member
// @Accept json
// @Produce json
// @Success 200 {object} MemberInfo
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/member/info [get]
func Info(c echo.Context) error {
	userID := c.Get("user_id").(int)

	var memberInfo MemberInfo
	query := `
        SELECT u.id, u.name, u.username, m.email, m.address, m.phone_num, m.payment_id, m.shipment_id
        FROM "user" u
        JOIN "member" m ON u.id = m.user_id
        WHERE u.id = $1
    `
	err := config.DB.QueryRow(query, userID).Scan(&memberInfo.ID, &memberInfo.Name, &memberInfo.Username, &memberInfo.Email, &memberInfo.Address, &memberInfo.PhoneNum, &memberInfo.PaymentID, &memberInfo.ShipmentID)
	if err != nil {
		log.Printf("Error fetching member info: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to fetch member info",
		})
	}

	return c.JSON(http.StatusOK, memberInfo)
}
