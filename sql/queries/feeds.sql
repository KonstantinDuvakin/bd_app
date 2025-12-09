-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, owner)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetAllFeeds :many
SELECT * FROM feeds;

-- name: GetNextFeedToFetch :many
SELECT * FROM feeds
ORDER BY last_fetched_at NULLS FIRST
LIMIT $1;

-- name: MarkFeedAsFetched :one
UPDATE feeds
SET last_fetched_at = NOW(),
updated_at = NOW()
WHERE id = $1
RETURNING *;
