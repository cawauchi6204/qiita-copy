CREATE TABLE
  IF NOT EXISTS stock_posts(
    id serial PRIMARY KEY,
    post_id VARCHAR(50),
    user_id VARCHAR(50),
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
  );
