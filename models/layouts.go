package models

import (
	"encoding/json"

	null "gopkg.in/volatiletech/null.v6"
)

type Layout struct {
	ID        int       `db:"id" json:"id"`
	CreatedAt null.Time `db:"created_at" json:"createdAt"`
	UpdatedAt null.Time `db:"updated_at" json:"updatedAt"`

	Module     string          `db:"module" json:"module"`
	LayoutType string          `db:"layout_type" json:"layoutType"`
	Meta       json.RawMessage `db:"meta" json:"meta"`
}

type CRMModule struct {
	Name  string `json:"name"`
	Label string `json:"label"`
}

type CRMField struct {
	Name  string                 `json:"name"`
	Label string                 `json:"label"`
	Type  string                 `json:"type"`
	Meta  map[string]interface{} `json:"meta"`
}

type Field struct {
	ID        int       `db:"id" json:"id"`
	CreatedAt null.Time `db:"created_at" json:"createdAt"`
	UpdatedAt null.Time `db:"updated_at" json:"updatedAt"`

	Module string          `db:"module" json:"module"`
	Name   string          `db:"name" json:"name"`
	Label  string          `db:"label" json:"label"`
	Type   string          `db:"type" json:"type"`
	Meta   json.RawMessage `db:"meta" json:"meta"`
}
