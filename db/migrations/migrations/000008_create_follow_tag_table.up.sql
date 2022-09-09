CREATE TABLE
  IF NOT EXISTS follow_tags(
    id serial PRIMARY KEY,
    tag_id VARCHAR(50) NOT NULL,
    user_id VARCHAR(50) NOT NULL
  );
