package repository

import (
	"product-service/src/internal/entity"

	"github.com/jmoiron/sqlx"
)

type ProductCategoryRepository interface {
	Insert(pc entity.ProductCategory) error
	Delete(id int64) error
	CountByID(id int64) (int64, error)
	FindAllProductsByCategoryID(page, size int, categoryID int64) ([]entity.ProductCategory, error)
}
type productCategoryRepository struct {
	db *sqlx.DB
}

func NewProductCategoryRepository(db *sqlx.DB) ProductCategoryRepository {
	return &productCategoryRepository{
		db: db,
	}
}
func (r *productCategoryRepository) Insert(pc entity.ProductCategory) error {
	query := "INSERT INTO product_categories (product_id, category_id) VALUES ($1, $2)"
	_, err := r.db.Exec(query, pc.ProductID, pc.CategoryID)
	if err != nil {
		return err
	}
	return nil
}

func (r *productCategoryRepository) Delete(id int64) error {
	query := `DELETE FROM product_categories WHERE id = $1`
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
func (r *productCategoryRepository) CountByID(id int64) (int64, error) {
	query := `SELECT COUNT(*) FROM product_categories WHERE id = $1`

	var count int64
	err := r.db.Get(&count, query, id)
	if err != nil {
		return 0, err
	}
	return count, nil
}
func (r *productCategoryRepository) FindAllProductsByCategoryID(page, size int, categoryID int64) ([]entity.ProductCategory, error) {
	offset := (page - 1) * size
	var productCategories []entity.ProductCategory
	query := "SELECT * FROM product_categories WHERE category_id = $1 LIMIT $2 OFFSET $3"
	err := r.db.Select(&productCategories, query, categoryID, size, offset)
	if err != nil {
		return nil, err
	}
	return productCategories, nil
}
