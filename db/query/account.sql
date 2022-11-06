-- name: CreateAccount :one
INSERT INTO accounts (
  user_id, 
  category_id, 
  title, 
  type, 
  description, 
  value, 
  date
  ) VALUES (
    $1, $2, $3, $4, $5, $6, $7
    ) RETURNING *;

-- name: GetAccount :one
SELECT * FROM accounts 
  WHERE id = $1 
  AND deleted_at IS NULL
  LIMIT 1;

-- name: GetAccounts :many
SELECT a.id, a.user_id, a.title, a.type, a.description, a.value, a.date, a.created_at, 
    c.title as category_title, c.type as category_type 
 FROM accounts a 
    LEFT JOIN categories c 
    ON c.id = a.category_id 
  WHERE a.user_id = $1 
    AND a.type = $2 
    AND a.category_id = $3 
    AND a.title like $4 
    AND a.description LIKE $5
    AND a.date = $6
    AND a.deleted_at IS NULL
  ORDER BY a.date DESC;

-- name: GetAllAccounts :many
  SELECT * from accounts 
    WHERE deleted_at IS NULL;


-- name: GetAccountsReports :one
SELECT SUM(value) AS sum_value FROM accounts
  WHERE user_id = $1 
    AND type = $2
    AND deleted_at IS NULL;

-- name: GetAccountsGraph :one
SELECT COUNT(*) FROM accounts
  WHERE user_id = $1 
    AND type = $2
    AND deleted_at IS NULL;

-- name: UpdateAccount :one
UPDATE accounts  
  SET title = $2, description = $3, value = $4, updated_at = CURRENT_TIMESTAMP
  WHERE id = $1 
  RETURNING *;

-- name: DeleteAccount :exec
UPDATE accounts 
  SET deleted_at = CURRENT_TIMESTAMP
  WHERE id = $1;

-- name: GetDeletedAccount :one
SELECT * FROM accounts 
  WHERE id = $1 
  AND deleted_at IS NOT NULL;

-- name: GetDeletedAccounts :many
SELECT * FROM accounts 
  WHERE deleted_at IS NOT NULL;
