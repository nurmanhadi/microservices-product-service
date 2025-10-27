package entity

type ProductCategory struct {
	ID         int64 `db:"id"`
	ProductID  int64 `db:"product_id"`
	CategoryID int   `db:"category_id"`
}
