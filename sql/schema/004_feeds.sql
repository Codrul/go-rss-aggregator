-- +goose UP

  CREATE TABLE feeds (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name VARCHAR(256) NOT NULL,
    url VARCHAR(256) NOT NULL UNIQUE,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE
  );


-- +goose DOWN
  DROP TABLE feeds;

