package models

import (
	"encoding/json"

	null "gopkg.in/volatiletech/null.v6"
)

type Products []Product

type Product struct {
	ID        int       `db:"id" json:"id"`
	CreatedAt null.Time `db:"created_at" json:"created_at"`
	UpdatedAt null.Time `db:"updated_at" json:"updated_at"`
	Source    string    `db:"source" json:"source"`
	SourceID  string    `db:"source_id" json:"source_id"`

	Name        string          `db:"name" json:"name"`
	Images      json.RawMessage `db:"images" json:"images"`
	SKU         string          `db:"sku" json:"sku"`
	Cost        float64         `db:"cost" json:"cost"`
	Price       float64         `db:"price" json:"price"`
	Description string          `db:"description" json:"description"`
}
