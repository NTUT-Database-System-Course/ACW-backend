package member

import (
    "database/sql"
)

type Member struct {
    ID       int    `json:"id"`
    Name     string `json:"name" validate:"required"`
    Password string `json:"password" validate:"required"`
    Username string `json:"username" validate:"required"`
    Address  string `json:"address"`
    Email    string `json:"email" validate:"required"`
    PhoneNum string `json:"phone_num"`
    PaymentID int   `json:"payment_id"`
    ShipmentID int  `json:"shipment_id"`
}

type RegistrationRequest struct {
    Name     string `json:"name" validate:"required"`
    Username string `json:"username" validate:"required"`
    Password string `json:"password" validate:"required"`
    Email    string `json:"email" validate:"required"`
}

type LoginRequest struct {
    Username string `json:"username" validate:"required"`
    Password string `json:"password" validate:"required"`
}

type MemberInfo struct {
    ID         int            `json:"id"`
    Name       string         `json:"name"`
    Username   string         `json:"username"`
    Email      string         `json:"email"`
    Address    sql.NullString `json:"address"`
    PhoneNum   sql.NullString `json:"phone_num"`
    PaymentID  sql.NullInt64  `json:"payment_id"`
    ShipmentID sql.NullInt64  `json:"shipment_id"`
}