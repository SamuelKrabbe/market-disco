-- name: CreateUser :one
INSERT INTO users (username, email, password_hash, cpf, role)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetUserByID :one
SELECT * FROM users WHERE id = $1 LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users ORDER BY id;

-- name: UpdateUser :one
UPDATE users
SET username = COALESCE($2, username),
    email    = COALESCE($3, email),
    role     = COALESCE($4, role)
WHERE id = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;

-- name: GetUserPoints :one
SELECT points FROM users WHERE id = $1;
