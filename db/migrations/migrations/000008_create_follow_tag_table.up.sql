CREATE TABLE
  IF NOT EXISTS follow_tags(
    id serial PRIMARY KEY,
    tag_id INTEGER NOT NULL,
    user_id INTEGER NOT NULL
  );
