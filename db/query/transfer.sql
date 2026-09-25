-- name: CreateTransfer :one
INSERT INTO transfers (
  from_account_id,
  to_account_id,
  amount
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetTransfer :one
SELECT * FROM transfers
WHERE id = $1 LIMIT 1;

-- name: GetTransfersByAccountID :many
SELECT id, from_account_id, to_account_id, amount, created_at
FROM (
    SELECT t.id, t.from_account_id, t.to_account_id, t.amount, t.created_at
    FROM transfers t
    WHERE t.from_account_id = $1
    UNION ALL
    SELECT t.id, t.from_account_id, t.to_account_id, t.amount, t.created_at
    FROM transfers t
    WHERE t.to_account_id = $1
) AS combined_transfers
ORDER BY created_at DESC, id DESC
LIMIT $2
OFFSET $3;
