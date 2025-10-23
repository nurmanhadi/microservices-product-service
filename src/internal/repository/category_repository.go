package repository

import (
	"product-service/src/internal/entity"

	"github.com/jmoiron/sqlx"
)

type CategoryRepository interface {
	InsertSubCategory(cat entity.Category) error
	InsertCategory(cat entity.Category) error
	CountByID(id int) (int, error)
	UpdateByID(cat entity.Category) error
	DeleteByID(id int) error
	FindAll() ([]entity.Category, error)
	FindAllSubCategoryByID(id int) ([]entity.Category, error)
	FindByID(id int) (*entity.Category, error)
}
type categoryRepository struct {
	db *sqlx.DB
}

func NewCategoryRepository(db *sqlx.DB) CategoryRepository {
	return &categoryRepository{
		db: db,
	}
}
func (r *categoryRepository) InsertCategory(cat entity.Category) error {
	query := `INSERT INTO categories (name) VALUES ($1)`
	_, err := r.db.Exec(query, cat.Name)
	if err != nil {
		return err
	}
	return nil
}
func (r *categoryRepository) InsertSubCategory(cat entity.Category) error {
	query := `INSERT INTO categories (name, parent_id) VALUES ($1, $2)`
	_, err := r.db.Exec(query, cat.Name, cat.ParentID)
	if err != nil {
		return err
	}
	return nil
}
func (r *categoryRepository) CountByID(id int) (int, error) {
	var count int
	query := `SELECT COUNT(*) FROM categories WHERE id = $1`
	err := r.db.Get(&count, query, id)
	if err != nil {
		return 0, err
	}
	return count, nil
}
func (r *categoryRepository) UpdateByID(cat entity.Category) error {
	query := `UPDATE categories SET name = $1 WHERE id = $2`
	_, err := r.db.Exec(query, cat.Name, cat.ID)
	return err
}
func (r *categoryRepository) DeleteByID(id int) error {
	query := `DELETE FROM categories WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}
func (r *categoryRepository) FindAll() ([]entity.Category, error) {
	var categories []entity.Category
	query := "select * from categories"
	err := r.db.Select(&categories, query)
	if err != nil {
		return nil, err
	}
	return categories, nil
}
func (r *categoryRepository) FindAllSubCategoryByID(id int) ([]entity.Category, error) {
	var categories []entity.Category
	query := "select * from categories where parent_id = $1"
	err := r.db.Select(&categories, query, id)
	if err != nil {
		return nil, err
	}
	return categories, nil
}
func (r *categoryRepository) FindByID(id int) (*entity.Category, error) {
	category := new(entity.Category)
	query := "select * from categories where id = $1"
	err := r.db.Get(category, query, id)
	if err != nil {
		return nil, err
	}
	return category, nil
}
