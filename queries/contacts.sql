-- name: create-contact
INSERT INTO contacts (
    first_name,
    last_name,
    email,
    phone,
    description,
    contact_type,
    status,
    score,
    subscriber_id,
    account_id,
    source,
    source_id
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12
) RETURNING id;

-- name: update-contact
UPDATE contacts
SET
    first_name = $2,
    last_name = $3,
    email = $4,
    phone = $5,
    description = $6,
    contact_type = $7,
    status = $8,
    score = $9,
    subscriber_id = $10,
    account_id = $11,
    source = $12,
    source_id = $13
WHERE id = $1;
