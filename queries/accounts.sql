-- name: create-account
INSERT INTO accounts (
    name,
    short_name,
    account_type,
    industry,
    sic_code,
    website,
    description,
    source,
    source_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id;

-- name: update-account
UPDATE accounts SET
    name = $2,
    short_name = $3,
    account_type = $4,
    industry = $5,
    sic_code = $6,
    website = $7,
    description = $8,
    source = $9,
    source_id = $10
WHERE id = $1;
