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

