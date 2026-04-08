package migrations

import (
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/knadh/koanf/v2"
	"github.com/knadh/stuffbin"
)

func CRMv1_0_1(db *sqlx.DB, fs stuffbin.FileSystem, ko *koanf.Koanf, lo *log.Logger) error {
	if _, err := db.Exec(`
		ALTER TABLE accounts ADD COLUMN deleted BOOLEAN DEFAULT FALSE;
		ALTER TABLE contacts ADD COLUMN deleted BOOLEAN DEFAULT FALSE;
		ALTER TABLE addresses ADD COLUMN deleted BOOLEAN DEFAULT FALSE;
		ALTER TABLE products ADD COLUMN deleted BOOLEAN DEFAULT FALSE;
		ALTER TABLE quotes ADD COLUMN deleted BOOLEAN DEFAULT FALSE;
		ALTER TABLE invoices ADD COLUMN deleted BOOLEAN DEFAULT FALSE;

		CREATE TABLE IF NOT EXISTS orders (
			id SERIAL PRIMARY KEY,
			created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
			updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
			source VARCHAR(255),
			source_id VARCHAR(255),
			deleted BOOLEAN DEFAULT FALSE,
			account_id INT NOT NULL REFERENCES accounts(id) ON DELETE CASCADE ON UPDATE CASCADE,
			contact_id INT NOT NULL REFERENCES contacts(id) ON DELETE CASCADE ON UPDATE CASCADE,
			name VARCHAR(255) NOT NULL UNIQUE,
			status VARCHAR(255) NOT NULL DEFAULT 'draft',
			due_date TIMESTAMP,
			description TEXT,
			total FLOAT,
			subtotal FLOAT,
			discount FLOAT,
			discount_type VARCHAR(255),
			tax FLOAT,
			tax_type VARCHAR(255)
		);
		CREATE INDEX IF NOT EXISTS idx_orders_source_source_id ON orders (source, source_id);
	`); err != nil {
		return err
	}
	return nil
}
