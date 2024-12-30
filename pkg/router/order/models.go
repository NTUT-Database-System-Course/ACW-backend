package order

import (
	"database/sql"
	"time"
)

type CreationRequest struct {
	Name           string `json:"name"`
	Description    string `json:"description"`
	State          string `json:"state"`
	Address        string `json:"address"`
	PhoneNum       string `json:"phone_num"`
	PaymentMethod  int    `json:"payment_method"`
	ShipmentMethod int    `json:"shipment_method"`
}

type Product struct {
	Name        string         `json:"name" validate:"required"`
	Description string         `json:"description" validate:"required"`
	Price       int            `json:"price" validate:"required"`
	VendorName  string         `json:"vendor_name" validate:"required"`
	BuildTime   time.Time      `json:"build_time"`
	ImageURL    sql.NullString `json:"image_url"`
	Count       int            `json:"count"`
}

type Order struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	Description    string    `json:"description"`
	State          string    `json:"state"`
	Address        string    `json:"address"`
	PhoneNum       string    `json:"phone_num"`
	PaymentMethod  int       `json:"payment_method"`
	ShipmentMethod int       `json:"shipment_method"`
	Products       []Product `json:"products"`
}
