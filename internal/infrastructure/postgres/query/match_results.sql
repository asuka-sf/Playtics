-- name: CreateMatchResult :one
INSERT INTO playtics.match_results (player_id, match_id, kill_count, death_count, score, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: GetLeaderboard :many
SELECT 
    p.id, 
    p.name, 
    SUM(mr.score) AS total_score, 
    SUM(mr.kill_count) AS total_kills, 
    SUM(mr.death_count) As total_deaths,
    RANK() OVER (ORDER BY SUM(mr.score) DESC) AS rank
FROM playtics.players p
JOIN playtics.match_results mr ON p.id = mr.player_id
GROUP BY p.id, p.name
ORDER BY SUM(mr.score) DESC;