CREATE TABLE
  IF NOT EXISTS likes(
    id serial PRIMARY KEY,
    like_user_id VARCHAR(50),
    post_id VARCHAR(255),
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
  );
