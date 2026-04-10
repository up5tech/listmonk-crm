package migrations

import (
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/knadh/koanf/v2"
	"github.com/knadh/stuffbin"
)

func CRMv1_0_2(db *sqlx.DB, fs stuffbin.FileSystem, ko *koanf.Koanf, lo *log.Logger) error {
	if _, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS layouts (
			id SERIAL PRIMARY KEY,
			created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
			updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
			module VARCHAR(255),
			layout_type VARCHAR(255),
			meta JSONB
		);

		CREATE INDEX IF NOT EXISTS idx_layouts_module_layout_type ON layouts (module, layout_type);

		UPDATE settings SET value = '["v6.1.3"]' WHERE key = 'migrations';
	`); err != nil {
		return err
	}
	return nil
}
