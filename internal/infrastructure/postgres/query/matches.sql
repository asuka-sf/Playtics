-- name: CreateMatch :one
INSERT INTO playtics.matches (id, duration_seconds, created_at)
VALUES ($1, $2, $3)
RETURNING *;