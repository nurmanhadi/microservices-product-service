package dto

type ProductCategoryAddRequest struct {
	ProductID  int64 `json:"product_id" validate:"required"`
	CategoryID int   `json:"category_id" validate:"required"`
}
