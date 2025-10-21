package repository

import (
	"fmt"
	"log"
	"product-service/src/dto"
	"product-service/src/internal/entity"

	"github.com/jmoiron/sqlx"
)

type ProductRepository interface {
	Insert(product entity.Product) error
	FindAll() ([]entity.Product, error)
	FindAllByBulkID(ids []int) ([]entity.Product, error)
	FindByID(id int64) (*entity.Product, error)
	CountByID(id int64) (int64, error)
	UpdateBulkQuantityByID(datas []dto.ProductUpdateBulkQuantity) error
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
func (r *productRepository) FindAllByBulkID(ids []int) ([]entity.Product, error) {
	products := []entity.Product{}
	query, args, err := sqlx.In("SELECT * FROM products WHERE id IN (?)", ids)
	if err != nil {
		return nil, err
	}
	query = r.db.Rebind(query)
	err = r.db.Select(&products, query, args...)
	if err != nil {
		return nil, err
	}
	return products, nil
}
func (r *productRepository) FindByID(id int64) (*entity.Product, error) {
	product := new(entity.Product)
	err := r.db.Get(product, "SELECT * FROM products WHERE id=$1", id)
	if err != nil {
		return nil, err
	}
	return product, nil
}
func (r *productRepository) CountByID(id int64) (int64, error) {
	var total int64
	err := r.db.Get(&total, "SELECT COUNT(*) FROM products WHERE id=$1", id)
	if err != nil {
		return 0, err
	}
	return total, nil
}
func (r *productRepository) UpdateBulkQuantityByID(datas []dto.ProductUpdateBulkQuantity) error {
	query1 := "UPDATE products SET quantity = CASE id "
	query2 := ""
	query3 := "END "
	query4 := "WHERE id IN (?)"
	ids := make([]int, 0, len(datas))
	for _, x := range datas {
		ids = append(ids, int(x.ID))
		query2 += fmt.Sprintf("WHEN %d THEN %d ", x.ID, x.Quantity)
	}
	finalQuery := query1 + query2 + query3 + query4
	query, args, err := sqlx.In(finalQuery, ids)
	if err != nil {
		return err
	}
	query = r.db.Rebind(query)
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	log.Println(query)
	_, err = tx.Exec(query, args...)
	if err != nil {
		err := tx.Rollback()
		if err != nil {
			return err
		}
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}
