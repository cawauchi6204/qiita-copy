CREATE TABLE
  IF NOT EXISTS stock_posts(
    id serial PRIMARY KEY,
    post_id INTEGER,
    user_id INTEGER,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
  );
