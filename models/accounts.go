package models

import null "gopkg.in/volatiletech/null.v6"

type Accounts []Account

type Account struct {
	ID        int         `db:"id" json:"id"`
	CreatedAt null.Time   `db:"created_at" json:"created_at"`
	UpdatedAt null.Time   `db:"updated_at" json:"updated_at"`
	Source    null.String `db:"source" json:"source"`
	SourceID  null.String `db:"source_id" json:"source_id"`

	Name        string `db:"name" json:"name"`
	ShortName   string `db:"short_name" json:"short_name"`
	AccountType string `db:"account_type" json:"account_type"`
	Industry    string `db:"industry" json:"industry"`
	SicCode     string `db:"sic_code" json:"sic_code"`
	Website     string `db:"website" json:"website"`
	Description string `db:"description" json:"description"`

	Address AddressRelationship `db:"address" json:"address"`
}
