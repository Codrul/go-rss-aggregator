-- +goose UP

  ALTER TABLE users 
  ADD CONSTRAINT unique_user_name UNIQUE(user_name);

-- +goose DOWN
  ALTER TABLE users
  DROP CONSTRAINT unique_user_name;
