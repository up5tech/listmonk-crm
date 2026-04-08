package migrations

import (
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/knadh/koanf/v2"
	"github.com/knadh/stuffbin"
)

func CRMv1_0_0(db *sqlx.DB, fs stuffbin.FileSystem, ko *koanf.Koanf, lo *log.Logger) error {
	if _, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS accounts (
			id SERIAL PRIMARY KEY,
			created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
			updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
			source VARCHAR(255),
			source_id VARCHAR(255),
			name VARCHAR(255) NOT NULL,
			short_name VARCHAR(255) NOT NULL,
			account_type VARCHAR(255) NOT NULL,
			industry VARCHAR(255),
			sic_code VARCHAR(255),
			website VARCHAR(255),
			description TEXT
		);
		CREATE INDEX IF NOT EXISTS idx_accounts_source ON accounts (source, source_id);

		CREATE TABLE IF NOT EXISTS contacts (
			id SERIAL PRIMARY KEY,
			created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
			updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
			source VARCHAR(255),
			source_id VARCHAR(255),
			contact_type VARCHAR(255) NOT NULL,
			first_name VARCHAR(255) NOT NULL,
			last_name VARCHAR(255) NOT NULL,
			email VARCHAR(255) NOT NULL,
			phone VARCHAR(255),
			description TEXT,
			status VARCHAR(255),
			score INT,
			birth_date TIMESTAMP,
			subscriber_id INT REFERENCES subscribers(id) ON DELETE CASCADE ON UPDATE CASCADE,
			account_id INT REFERENCES accounts(id) ON DELETE CASCADE ON UPDATE CASCADE
		);
		CREATE INDEX IF NOT EXISTS idx_contacts_source ON contacts (source, source_id);

		CREATE TABLE IF NOT EXISTS addresses (
			id SERIAL PRIMARY KEY,
			created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
			updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
			source VARCHAR(255),
			source_id VARCHAR(255),
			address VARCHAR(255) NOT NULL,
			city VARCHAR(255),
			zip VARCHAR(255),
			country VARCHAR(255)
		);
		CREATE INDEX IF NOT EXISTS idx_addresses_source ON addresses (source, source_id);

		CREATE TABLE IF NOT EXISTS address_relationships (
			id SERIAL PRIMARY KEY,
			created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
			updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
			address_id INT NOT NULL REFERENCES addresses(id) ON DELETE CASCADE ON UPDATE CASCADE,
			relation_id INT NOT NULL,
			relation_type VARCHAR(255) NOT NULL
		);
		CREATE INDEX IF NOT EXISTS idx_address_relationships_relation ON address_relationships (relation_id, relation_type);

		CREATE TABLE IF NOT EXISTS products (
			id SERIAL PRIMARY KEY,
			created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
			updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
			source VARCHAR(255),
			source_id VARCHAR(255),
			name VARCHAR(255) NOT NULL,
			sku VARCHAR(255) UNIQUE,
			images JSONB NOT NULL DEFAULT '{}',
			description TEXT,
			cost FLOAT,
			price FLOAT NOT NULL
		);
		CREATE INDEX IF NOT EXISTS idx_products_source ON products (source, source_id);

		CREATE TABLE IF NOT EXISTS quotes (
			id SERIAL PRIMARY KEY,
			created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
			updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
			source VARCHAR(255),
			source_id VARCHAR(255),
			contact_id INT NOT NULL REFERENCES contacts(id) ON DELETE CASCADE ON UPDATE CASCADE,
			account_id INT REFERENCES accounts(id) ON DELETE CASCADE ON UPDATE CASCADE,
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
		CREATE INDEX IF NOT EXISTS idx_quotes_source ON quotes (source, source_id);

		CREATE TABLE IF NOT EXISTS invoices (
			id SERIAL PRIMARY KEY,
			created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
			updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
			source VARCHAR(255),
			source_id VARCHAR(255),
			contact_id INT NOT NULL REFERENCES contacts(id) ON DELETE CASCADE ON UPDATE CASCADE,
			account_id INT REFERENCES accounts(id) ON DELETE CASCADE ON UPDATE CASCADE,
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
		CREATE INDEX IF NOT EXISTS idx_invoices_source ON invoices (source, source_id);

		CREATE TABLE IF NOT EXISTS line_items (
			id SERIAL PRIMARY KEY,
			created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
			updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
			quote_id INT REFERENCES quotes(id) ON DELETE CASCADE ON UPDATE CASCADE,
			invoice_id INT REFERENCES invoices(id) ON DELETE CASCADE ON UPDATE CASCADE,
			product_id INT REFERENCES products(id) ON DELETE CASCADE ON UPDATE CASCADE,
			item_type VARCHAR(255) NOT NULL,
			name VARCHAR(255) NOT NULL,
			quantity INT NOT NULL,
			price FLOAT NOT NULL,
			description TEXT,
			images JSONB NOT NULL DEFAULT '{}'
		);
	`); err != nil {
		print(err.Error())
		return err
	}
	return nil
}
