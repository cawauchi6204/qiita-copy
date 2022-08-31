CREATE TABLE
  IF NOT EXISTS user_follows(
    id serial PRIMARY KEY,
    follow_user_id INTEGER,
    followed_user_id INTEGER,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
  );
