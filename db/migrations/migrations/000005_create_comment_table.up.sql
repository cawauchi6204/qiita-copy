CREATE TABLE
  IF NOT EXISTS comments(
    id serial PRIMARY KEY,
    content TEXT NOT NULL,
    post_id INTEGER NOT NULL,
    commented_by INTEGER NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
  );
