-- name: GetUserById :one
SELECT *
FROM users
WHERE id = $1
LIMIT 1;
-- name: GetAllUsers :many
SELECT *
FROM users
ORDER BY id ASC;
-- name: CreateUser :one
INSERT INTO users (username, email, password)
VALUES ($1, $ 2, $ 3)
RETURNING *;
-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;
-- name: GetUserByEmail :one
SELECT *
FROM users
WHERE email = $1
LIMIT 1;
-- name: GetUserByUsername :one
SELECT *
FROM users
WHERE username = $1
LIMIT 1;
-- name: UpdateUserById :one
UPDATE users
SET