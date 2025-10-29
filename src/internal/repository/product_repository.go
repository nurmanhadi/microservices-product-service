package repository

import (
	"fmt"
	"product-service/src/dto"
	"product-service/src/internal/entity"
	"strings"

	"github.com/jmoiron/sqlx"
)

type ProductRepository interface {
	Insert(product entity.Product) (int64, error)
	FindAll() ([]entity.Product, error)
	FindAllByBulkID(ids []int) ([]entity.Product, error)
	FindByID(id int64) (*entity.Product, error)
	CountByID(id int64) (int64, error)
	UpdateBulkQuantityByID(datas []dto.ProductUpdateBulkQuantity) error
	UpdateProduct(id int64, data dto.ProductUpdateRequest) error
	FindInID(ids []int64) ([]entity.Product, error)
}
type productRepository struct {
	db *sqlx.DB
}

func NewProductRepository(db *sqlx.DB) ProductRepository {
	return &productRepository{
		db: db,
	}
}
func (r *productRepository) Insert(product entity.Product) (int64, error) {
	var id int64
	err := r.db.QueryRow("INSERT INTO products(name, quantity, price) VALUES($1, $2, $3) RETURNING id", product.Name, product.Quantity, product.Price).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
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
func (r *productRepository) UpdateProduct(id int64, data dto.ProductUpdateRequest) error {
	setClaus := []string{}
	args := []interface{}{}
	argPos := 1
	if data.Name != nil {
		setClaus = append(setClaus, fmt.Sprintf("name = $%d", argPos))
		args = append(args, data.Name)
		argPos++
	}
	if data.Price != nil {
		setClaus = append(setClaus, fmt.Sprintf("price = $%d", argPos))
		args = append(args, data.Price)
		argPos++
	}
	if data.Quantity != nil {
		setClaus = append(setClaus, fmt.Sprintf("quantity = $%d", argPos))
		args = append(args, data.Quantity)
		argPos++
	}
	if len(setClaus) == 0 {
		return nil
	}
	query := fmt.Sprintf("UPDATE products SET %s WHERE id = $%d", strings.Join(setClaus, ", "), argPos)
	args = append(args, id)
	_, err := r.db.Exec(query, args...)
	if err != nil {
		return nil
	}
	return nil
}
func (r *productRepository) FindInID(ids []int64) ([]entity.Product, error) {
	query := "SELECT * FROM products WHERE id IN(?)"
	finalQuery, args, err := sqlx.In(query, ids)
	if err != nil {
		return nil, err
	}
	endQuery := r.db.Rebind(finalQuery)
	var products []entity.Product
	err = r.db.Select(&products, endQuery, args...)
	if err != nil {
		return nil, err
	}
	return products, err
}
