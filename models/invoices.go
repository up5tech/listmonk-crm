package models

const (
	InvoiceStatusDraft   = "draft"
	InvoiceStatusPaid    = "paid"
	InvoiceStatusPending = "pending"
)

type Invoices []Invoice

type Invoice struct {
	ID        int    `db:"id" json:"id"`
	CreatedAt string `db:"created_at" json:"created_at"`
	UpdatedAt string `db:"updated_at" json:"updated_at"`

	ContactID int `db:"contact_id" json:"contact_id"`
	AccountID int `db:"account_id" json:"account_id"`

	Name         string  `db:"name" json:"name"`
	Status       string  `db:"status" json:"status"`
	DueDate      string  `db:"due_date" json:"due_date"`
	Description  string  `db:"description" json:"description"`
	Total        float64 `db:"total" json:"total"`
	SubTotal     float64 `db:"sub_total" json:"sub_total"`
	Discount     float64 `db:"discount" json:"discount"`
	DiscountType string  `db:"discount_type" json:"discount_type"`
	Tax          float64 `db:"tax" json:"tax"`
	TaxType      string  `db:"tax_type" json:"tax_type"`

	Address AddressRelationship `json:"address"`
	Items   LineItems           `json:"items"`
}
