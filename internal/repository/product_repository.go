package repository

import (
	"product-service/internal/entity"

	"github.com/jmoiron/sqlx"
)

type ProductRepository interface {
	Insert(product entity.Product) error
	FindAll() ([]entity.Product, error)
	FindByID(id int64) (*entity.Product, error)
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
