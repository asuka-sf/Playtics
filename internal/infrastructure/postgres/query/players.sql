-- name: CreatePlayer :one
INSERT INTO playtics.players (id, name, email, image_url)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetPlayer :one
SELECT * FROM playtics.players
WHERE id = $1;