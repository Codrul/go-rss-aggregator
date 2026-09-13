-- +goose UP

  CREATE TABLE users (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    user_name VARCHAR(256) NOT NULL
  );


-- +goose DOWN
  DROP TABLE users;
