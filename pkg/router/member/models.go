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
}

type RegistrationRequest struct {
	Name     string `json:"name" validate:"required"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Email    string `json:"email" validate:"required"`
}

type UpdateRequest struct {
	Name     string `json:"name"`
	Address  string `json:"address"`
	Email    string `json:"email"`
	PhoneNum string `json:"phone_num"`
}

type MemberInfo struct {
	ID       int            `json:"id"`
	Name     string         `json:"name"`
	Username string         `json:"username"`
	Email    string         `json:"email"`
	Address  sql.NullString `json:"address"`
	PhoneNum sql.NullString `json:"phone_num"`
}
