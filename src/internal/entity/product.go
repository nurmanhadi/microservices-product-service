package entity

import "time"

type Product struct {
	ID        int64     `db:"id"`
	Name      string    `db:"name"`
	Quantity  int       `db:"quantity"`
	Price     int64     `db:"price"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
