-- +goose Up
CREATE TABLE posts
(
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    title VARCHAR(500) NOT NULL,
    url VARCHAR(1000) UNIQUE NOT NULL,
    description VARCHAR(10000) NOT NULL,
    published_at TIMESTAMP NOT NULL,
    feed_id UUID NOT NULL,
    FOREIGN KEY (feed_id) REFERENCES feeds(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE posts;