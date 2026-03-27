-- name: CreatePlayer :one
INSERT INTO playtics.players (id, name, email, image_url, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetPlayer :one
SELECT * FROM playtics.players
WHERE id = $1;