package models

import null "gopkg.in/volatiletech/null.v6"

type Addresses []Address

type Address struct {
	ID        int       `db:"id" json:"id"`
	CreatedAt null.Time `db:"created_at" json:"created_at"`
	UpdatedAt null.Time `db:"updated_at" json:"updated_at"`

	AddressLine1 null.String `db:"address_line1" json:"address_line1"`
	AddressLine2 null.String `db:"address_line2" json:"address_line2"`
	City         null.String `db:"city" json:"city"`
	State        null.String `db:"state" json:"state"`
	Zip          null.String `db:"zip" json:"zip"`
	Country      null.String `db:"country" json:"country"`
}

type AddressRelationship struct {
	ID        int       `db:"id" json:"id"`
	CreatedAt null.Time `db:"created_at" json:"created_at"`
	UpdatedAt null.Time `db:"updated_at" json:"updated_at"`

	AddressID int  `db:"address_id" json:"address_id"`
	IsPrimary bool `db:"is_primary" json:"is_primary"`

	RelationID   int    `db:"relation_id" json:"relation_id"`
	RelationType string `db:"relation_type" json:"relation_type"`
}
