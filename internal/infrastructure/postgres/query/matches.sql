-- name: CreateMatch :one
INSERT INTO playtics.matches (id, duration_seconds)
VALUES ($1, $2)
RETURNING *;