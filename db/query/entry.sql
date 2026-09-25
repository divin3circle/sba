-- name: CreateEntry :one
INSERT INTO entries (
  account_id,
  amount
) VALUES ( 
  $1, $2
) RETURNING *;

-- name: GetEntry :one
SELECT * from entries
WHERE id = $1 LIMIT 1;

-- name: GetEntriesByAccountID :many
SELECT * from entries
WHERE account_id = $1
ORDER BY created_at DESC
LIMIT $2
OFFSET $3;
