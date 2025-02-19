-- name: CreateFeedFollow :one
WITH inserted_feed_follow AS (
    INSERT INTO feed_follow (created_at, updated_at, user_id, feed_id)
    VALUES (
        $1,
        $2,
        $3,
        $4
        )
    RETURNING *
)
SELECT iff.*, f.name AS feed_name, u.name AS user_name
FROM inserted_feed_follow iff
INNER JOIN users u
ON u.id = iff.user_id
INNER JOIN feeds f
ON f.id = iff.feed_id;

-- name: GetFeedFollowsForUser :many
SELECT ff.*, f.name AS feed_name, u.name AS user_name
FROM feed_follow ff
INNER JOIN feeds f
ON ff.feed_id = f.id
INNER JOIN users u
ON ff.user_id = u.id
WHERE u.name = $1;

-- name: DeleteFeedFolow :exec
DELETE
FROM feed_follow ff
USING feeds f, users u
WHERE u.id = ff.user_id AND f.id = ff.feed_id AND u.name = sqlc.arg(user_name) AND f.url = sqlc.arg(feed_url);