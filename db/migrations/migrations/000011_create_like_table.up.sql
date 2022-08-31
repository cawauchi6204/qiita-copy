CREATE TABLE
  IF NOT EXISTS likes(
    id serial PRIMARY KEY,
    like_user_id INTEGER,
    liked_user_id INTEGER,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
  );
