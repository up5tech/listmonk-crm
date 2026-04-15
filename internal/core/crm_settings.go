package core

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"unicode"

	"github.com/knadh/listmonk/models"
)

func (c *Core) GetModules() []models.CRMModule {
	var modules = []models.CRMModule{
		{
			Name:  "Accounts",
			Label: "Accounts",
		},
		{
			Name:  "Contacts",
			Label: "Contacts",
		},
		{
			Name:  "Deals",
			Label: "Deals",
		},
		{
			Name:  "Tasks",
			Label: "Tasks",
		},
		{
			Name:  "Notes",
			Label: "Notes",
		},
		{
			Name:  "Products",
			Label: "Products",
		},
		{
			Name:  "Quotes",
			Label: "Quotes",
		},
		{
			Name:  "Orders",
			Label: "Orders",
		},
		{
			Name:  "Invoices",
			Label: "Invoices",
		},
	}
	return modules
}

func (c *Core) ModuleRegistry() map[string]interface{} {
	var MapRegistry = map[string]interface{}{
		"Layout":  models.Layout{},
		"Account": models.Account{},
		"Contact": models.Contact{},
		"Product": models.Product{},
		"Quote":   models.Quote{},
		"Invoice": models.Invoice{},
		"Address": models.Address{},
	}
	return MapRegistry
}

func (c *Core) MapGoType(t reflect.Type) string {
	switch t.String() {
	case "int", "int64":
		return "number"
	case "string":
		return "string"
	case "json.RawMessage":
		return "json"
	case "null.Time":
		return "datetime"
	default:
		return "unknown"
	}
}

func (c *Core) ToLabel(name string) string {
	// ID -> ID, CreatedAt -> Created At
	var result []rune
	for i, r := range name {
		if i > 0 && unicode.IsUpper(r) {
			result = append(result, ' ')
		}
		result = append(result, r)
	}
	return string(result)
}

func (c *Core) ParseRawMeta(tag string) map[string]interface{} {
	meta := map[string]interface{}{}

	if tag == "" {
		return meta
	}

	// ví dụ: "ui:json,readonly:true"
	parts := strings.Split(tag, ",")
	for _, p := range parts {
		kv := strings.SplitN(p, ":", 2)
		if len(kv) == 2 {
			meta[kv[0]] = kv[1]
		}
	}

	return meta
}

func (c *Core) ExtractFields(model interface{}) []models.CRMField {
	t := reflect.TypeOf(model)

	// nếu là pointer thì lấy Elem
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	var fields []models.CRMField

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)

		dbTag := f.Tag.Get("db")
		if dbTag == "" || dbTag == "-" {
			continue
		}

		// jsonTag := f.Tag.Get("json")
		// label := f.Tag.Get("label")
		// metaTag := f.Tag.Get("meta")
		label := ""

		if label == "" {
			label = c.ToLabel(f.Name)
		}

		fieldType := c.MapGoType(f.Type)
		meta := map[string]interface{}{}

		fields = append(fields, models.CRMField{
			Name:  dbTag,
			Label: label,
			Type:  fieldType,
			Meta:  meta,
		})
	}

	return fields
}

func (c *Core) GetModuleFields(module string) ([]models.CRMField, error) {
	model, ok := c.ModuleRegistry()[module]
	if !ok {
		return nil, fmt.Errorf("model not found: %s", module)
	}

	fields := c.ExtractFields(model)

	return fields, nil
}

func (c *Core) SaveField(module string, field models.CRMField) error {
	meta, _ := json.Marshal(field.Meta)
	saveField := models.Field{
		Module: module,
		Name:   field.Name,
		Label:  field.Label,
		Type:   field.Type,
		Meta:   meta,
	}
	// check field exist
	var fieldExist models.Field
	if err := c.db.Get(&fieldExist, "SELECT id FROM fields WHERE module = $1 AND name = $2", module, field.Name); err != nil {
		return err
	}

	// insert field if not exist
	if fieldExist.ID == 0 {
		_, err := c.q.CreateField.Exec(saveField)
		if err != nil {
			return err
		}
		return nil
	}

	// update field if exist
	_, err := c.q.UpdateField.Exec(saveField.ID, saveField.Module, saveField.Name, saveField.Label, saveField.Type, saveField.Meta)
	if err != nil {
		return err
	}
	return nil
}

func (c *Core) AutoImportFields() error {
	modules := c.GetModules()
	for _, module := range modules {
		fields, err := c.GetModuleFields(module.Name)
		if err != nil {
			return err
		}
		for _, field := range fields {
			if err := c.SaveField(module.Name, field); err != nil {
				return err
			}
		}
	}

	return nil
}
