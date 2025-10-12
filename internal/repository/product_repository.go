package repository

import (
	"product-service/internal/entity"

	"github.com/jmoiron/sqlx"
)

type ProductRepository interface {
	Insert(product entity.Product) error
	FindAll() ([]entity.Product, error)
	FindByID(id int64) (*entity.Product, error)
	UpdateQuantity(id, quantity int64) error
	CountByID(id int64) (int64, error)
}
type productRepository struct {
	db *sqlx.DB
}

func NewProductRepository(db *sqlx.DB) ProductRepository {
	return &productRepository{
		db: db,
	}
}
func (r *productRepository) Insert(product entity.Product) error {
	_, err := r.db.Exec("INSERT INTO products(name, quantity, price) VALUES($1, $2, $3)", product.Name, product.Quantity, product.Price)
	if err != nil {
		return err
	}
	return nil
}
func (r *productRepository) FindAll() ([]entity.Product, error) {
	products := []entity.Product{}
	err := r.db.Select(&products, "SELECT * FROM products")
	if err != nil {
		return nil, err
	}
	return products, nil
}
func (r *productRepository) FindByID(id int64) (*entity.Product, error) {
	product := entity.Product{}
	err := r.db.Get(&product, "SELECT * FROM products WHERE id=$1", id)
	if err != nil {
		return nil, err
	}
	return &product, nil
}
func (r *productRepository) CountByID(id int64) (int64, error) {
	var total int64
	err := r.db.Get(&total, "SELECT COUNT(*) FROM products WHERE id=$1", id)
	if err != nil {
		return 0, err
	}
	return total, nil
}
func (r *productRepository) UpdateQuantity(id, quantity int64) error {
	_, err := r.db.Exec("UPDATE products SET quantity = $1 WHERE id = $2", quantity, id)
	if err != nil {
		return err
	}
	return nil
}
