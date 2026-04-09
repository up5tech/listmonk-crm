package models

import null "gopkg.in/volatiletech/null.v6"

type Accounts []Account

type Account struct {
	ID        int         `db:"id" json:"id"`
	CreatedAt null.Time   `db:"created_at" json:"createdAt"`
	UpdatedAt null.Time   `db:"updated_at" json:"updatedAt"`
	Source    null.String `db:"source" json:"source"`
	SourceID  null.String `db:"source_id" json:"sourceID"`

	Name        string `db:"name" json:"name"`
	ShortName   string `db:"short_name" json:"shortName"`
	AccountType string `db:"account_type" json:"accountType"`
	Industry    string `db:"industry" json:"industry"`
	SicCode     string `db:"sic_code" json:"sicCode"`
	Website     string `db:"website" json:"website"`
	Description string `db:"description" json:"description"`

	Address AddressRelationship `db:"address" json:"address"`
}
