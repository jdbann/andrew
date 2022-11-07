-- name: CreatePost :one
INSERT INTO posts (slug, title, summary, body)
    VALUES ($1, $2, $3, $4)
RETURNING
    *;

-- name: GetPosts :many
SELECT
    *
FROM
    posts;

