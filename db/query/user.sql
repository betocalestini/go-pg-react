-- name: CreateUser :one
INSERT INTO users (
  name, 
  email, 
  password
  ) VALUES (
    $1, $2, $3
    ) RETURNING *;

-- name: GetUser :one
SELECT * FROM users 
  WHERE email = $1 
  AND deleted_at IS NULL
  LIMIT 1;

-- name: GetUserById :one
SELECT * FROM users 
  WHERE id = $1 
  AND deleted_at IS NULL
  LIMIT 1;

-- name: GetUsers :many
SELECT * FROM users 
  WHERE deleted_at IS NULL;

-- name: DeleteUser :exec
UPDATE users SET deleted_at = CURRENT_TIMESTAMP
  WHERE id = $1;

  -- name: GetDeletedUsers :many
SELECT * FROM users 
  WHERE deleted_at IS NOT NULL;

  -- name: GetDeletedUser :one
SELECT * FROM users 
  WHERE deleted_at IS NOT NULL
  AND id = $1;
