package models

import null "gopkg.in/volatiletech/null.v6"

const (
	ContactStatusActive   = "active"
	ContactStatusInactive = "inactive"
	ContactStatusPending  = "pending"

	ContactTypeCustomer = "Customer"
	ContactTypeLead     = "Lead"
)

type Contacts []Contact

type Contact struct {
	ID        int         `db:"id" json:"id"`
	CreatedAt null.Time   `db:"created_at" json:"created_at"`
	UpdatedAt null.Time   `db:"updated_at" json:"updated_at"`
	Source    null.String `db:"source" json:"source"`
	SourceID  null.String `db:"source_id" json:"source_id"`

	ContactType  string   `db:"contact_type" json:"contact_type"`
	SubscriberID null.Int `db:"subscriber_id" json:"subscriber_id"`
	AccountID    null.Int `db:"account_id" json:"account_id"`

	FirstName   null.String `db:"first_name" json:"first_name"`
	LastName    string      `db:"last_name" json:"last_name"`
	Email       string      `db:"email" json:"email"`
	Phone       null.String `db:"phone" json:"phone"`
	Description string      `db:"description" json:"description"`

	Status string   `db:"status" json:"status"`
	Score  null.Int `db:"score" json:"score"`

	Addresses []AddressRelationship `db:"addresses" json:"addresses"`
}
