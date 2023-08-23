-- name: CreateUser :exec
INSERT INTO users (
  id, username, password, role, is_active
) VALUES (
  $1, $2, $3, $4, $5
);

-- name: ListUser :many
SELECT id, username, role, is_active
FROM users;

-- name: GetUser :one
SELECT id, username, role, is_active
FROM users
WHERE id = $1;

-- name: GetUsernamePassword :one
SELECT id, username, password, role, is_active
FROM users
WHERE username = $1;

-- name: UpdateUser :exec
UPDATE users
SET username = $2, password = $3, role = $4, is_active = $5
WHERE id = $1;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;