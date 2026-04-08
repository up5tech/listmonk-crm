package models

import (
	"encoding/json"

	null "gopkg.in/volatiletech/null.v6"
)

const (
	QuoteStatusDraft   = "draft"
	QuoteStatusActive  = "active"
	QuoteStatusPaid    = "paid"
	QuoteStatusExpired = "expired"
)

type Quotes []Quote

type Quote struct {
	ID        int       `db:"id" json:"id"`
	CreatedAt null.Time `db:"created_at" json:"created_at"`
	UpdatedAt null.Time `db:"updated_at" json:"updated_at"`
	Source    string    `db:"source" json:"source"`
	SourceID  string    `db:"source_id" json:"source_id"`

	Name      string   `db:"name" json:"name"`
	ContactID int      `db:"contact_id" json:"contact_id"`
	AccountID null.Int `db:"account_id" json:"account_id"`

	Address AddressRelationship `json:"address"`
	Items   LineItems           `json:"items"`

	Status       string    `db:"status" json:"status"`
	DueDate      null.Time `db:"due_date" json:"due_date"`
	Description  string    `db:"description" json:"description"`
	Total        float64   `db:"total" json:"total"`
	Subtotal     float64   `db:"subtotal" json:"subtotal"`
	Discount     float64   `db:"discount" json:"discount"`
	DiscountType string    `db:"discount_type" json:"discount_type"`
	Tax          float64   `db:"tax" json:"tax"`
	TaxType      string    `db:"tax_type" json:"tax_type"`
}

type LineItems []LineItem

type LineItem struct {
	ID        int       `db:"id" json:"id"`
	CreatedAt null.Time `db:"created_at" json:"created_at"`
	UpdatedAt null.Time `db:"updated_at" json:"updated_at"`

	QuoteID   null.Int `db:"quote_id" json:"quote_id"`
	InvoiceID null.Int `db:"invoice_id" json:"invoice_id"`

	ItemType    string          `db:"item_type" json:"item_type"`
	Name        string          `db:"name" json:"name"`
	Quantity    int             `db:"quantity" json:"quantity"`
	Price       float64         `db:"price" json:"price"`
	Description string          `db:"description" json:"description"`
	Images      json.RawMessage `db:"images" json:"images"`

	ProductID null.Int `db:"product_id" json:"product_id"`
}
