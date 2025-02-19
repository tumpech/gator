-- name: CreateFeed :one
INSERT INTO feeds (created_at, updated_at, name, url, user_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5
)
RETURNING *;

-- name: ListFeeds :many
SELECT *
FROM feeds;

-- name: GetFeedByUrl :one
SELECT *
FROM feeds
WHERE url = $1;