package migrations

import (
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/knadh/koanf/v2"
	"github.com/knadh/stuffbin"
)

func CRMv1_0_3(db *sqlx.DB, fs stuffbin.FileSystem, ko *koanf.Koanf, lo *log.Logger) error {
	if _, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS fields (
			id SERIAL PRIMARY KEY,
			created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
			updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
			module VARCHAR(255) NOT NULL,
			name VARCHAR(255) NOT NULL,
			label VARCHAR(255) NOT NULL,
			type VARCHAR(255) NOT NULL,
			meta JSONB
		);
		CREATE UNIQUE INDEX IF NOT EXISTS idx_fields_module_name ON fields (module, name);

		UPDATE settings SET value = '["v6.1.4"]' WHERE key = 'migrations';
	`); err != nil {
		return err
	}
	return nil
}
