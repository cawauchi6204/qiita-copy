CREATE TABLE
  IF NOT EXISTS comments(
    id serial PRIMARY KEY,
    content TEXT NOT NULL,
    post_id VARCHAR(255) NOT NULL,
    commented_by VARCHAR(50) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
  );
