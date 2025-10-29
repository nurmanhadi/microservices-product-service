package collection

type ProductCollection struct {
	ID         string   `json:"id"`
	Name       string   `json:"name"`
	Quantity   int      `json:"quantity"`
	Price      int64    `json:"price"`
	CreatedAt  int64    `json:"created_at"`
	UpdatedAt  int64    `json:"updated_at"`
	Categories []string `json:"categories"`
}
