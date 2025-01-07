-- name: CreatePost :one
INSERT INTO posts (id, created_at, updated_at, title, url, description, published_at, feed_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8
)
RETURNING *;

-- name: GetPostsForUser :many
SELECT * FROM posts
INNER JOIN feed_follows ON posts.feed_id = feed_follows.feed_id
WHERE $1 = feed_follows.user_id
ORDER BY published_at DESC NULLS LAST
LIMIT $2;