CREATE TABLE posts (
    id bigserial PRIMARY KEY,
    slug text NOT NULL UNIQUE,
    title text NOT NULL,
    summary text NOT NULL,
    body text NOT NULL,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

