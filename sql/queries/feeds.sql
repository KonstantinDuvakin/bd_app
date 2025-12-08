-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, owner)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetAllFeeds :many
SELECT * FROM feeds;