package dto

import "time"

type ProductAddRequest struct {
	Name     string `json:"name" validate:"required,max=100"`
	Quantity int    `json:"quantity" validate:"required"`
	Price    int64  `json:"price" validate:"required"`
}
type ProductUpdateRequest struct {
	Name     *string `json:"name" validate:"omitempty,max=100"`
	Quantity *int    `json:"quantity" validate:"omitempty"`
	Price    *int64  `json:"price" validate:"omitempty"`
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
