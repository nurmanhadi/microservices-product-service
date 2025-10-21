package dto

import "time"

type ProductAddRequest struct {
	Name     string `json:"name" validate:"required,max=100"`
	Quantity int    `json:"quantity" validate:"required"`
	Price    int64  `json:"price" validate:"required"`
}
type ProductResponse struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Quantity  int       `json:"quantity"`
	Price     int64     `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
type ProductUpdateBulkQuantity struct {
	ID       int64 `json:"id"`
	Quantity int   `json:"quantity"`
}
