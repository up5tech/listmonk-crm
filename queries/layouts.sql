-- name: create-field
INSERT INTO fields (
    module,
    name,
    label,
    type,
    meta) VALUES ($1, $2, $3, $4, $5) RETURNING id;

-- name: update-field
UPDATE fields SET
    module = $2,
    name = $3,
    label = $4,
    type = $5,
    meta = $6
WHERE id = $1;
