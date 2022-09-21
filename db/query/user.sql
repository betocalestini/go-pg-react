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