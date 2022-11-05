-- name: CreateCategory :one
INSERT INTO categories (
  user_id, 
  title, 
  type, 
  description
  ) VALUES (
    $1, $2, $3, $4
    ) RETURNING *;

-- name: GetCategoryById :one
SELECT * FROM categories 
  WHERE id = $1 
  AND deleted_at IS NULL
  LIMIT 1;

-- name: GetCategoriesByTitle :many
SELECT * FROM categories 
  WHERE title = $1 
  AND deleted_at IS NULL;


-- name: GetCategories :many
SELECT * FROM categories 
  WHERE user_id = $1
  AND deleted_at IS NULL
  AND type = $2 
  AND title like $3 
  AND description LIKE $4;


-- name: GetAllCategories :many
SELECT * FROM categories 
  WHERE deleted_at IS NULL;


-- name: UpdateCategory :one
UPDATE categories 
  SET title = $2, description = $3, type = $4, updated_at = CURRENT_TIMESTAMP
  WHERE id = $1 
  RETURNING *;

-- name: DeleteCategory :exec
UPDATE categories 
  SET deleted_at = CURRENT_TIMESTAMP
  WHERE id = $1;

-- name: GetDeletedCategories :many
SELECT * FROM categories 
  WHERE deleted_at IS NOT NULL;
  
-- name: GetDeletedCategory :one
SELECT * FROM categories 
  WHERE deleted_at IS NOT NULL
  AND id = $1;
