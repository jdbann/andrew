-- name: CreatePost :one
INSERT INTO posts (slug, title, summary, body)
    VALUES ($1, $2, $3, $4)
RETURNING
    *;

-- name: GetPost :one
SELECT
    *
FROM
    posts
WHERE
    slug = $1
LIMIT 1;

-- name: GetPosts :many
SELECT
    *
FROM
    posts;

-- name: SeedPost :exec
INSERT INTO posts (slug, title, summary, body, created_at)
    VALUES ($1, $2, $3, $4, $5)
ON CONFLICT (slug)
    DO UPDATE SET
        title = EXCLUDED.title, summary = EXCLUDED.summary, body = EXCLUDED.body, created_at = EXCLUDED.created_at;

