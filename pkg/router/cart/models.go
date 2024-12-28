package cart

import (
    "database/sql"
    "time"
)

type CartItem struct {
    ID           int            `json:"id"`
    Name         string         `json:"name" validate:"required"`
    Description  string         `json:"description" validate:"required"`
    Price        int            `json:"price" validate:"required"`
    VendorID     int            `json:"vendor_id" validate:"required"`
    Remain       int            `json:"remain" validate:"required"`
    Disability   bool           `json:"disability" validate:"required"`
    BuildTime    time.Time      `json:"build_time"`
    ImageURL     sql.NullString `json:"image_url"`
    Tags         []int          `json:"tags"`
    Count        int            `json:"count"`
    CreationTime time.Time      `json:"creation_time"`
}