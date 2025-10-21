package dto

type OrderConsumerResponse struct {
	OrderID   string `json:"order_id"`
	ProductID int64  `json:"product_id"`
	Quantity  int    `json:"quantity"`
}
