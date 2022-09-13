CREATE TABLE
  IF NOT EXISTS post_tags(
    id serial PRIMARY KEY,
    post_id VARCHAR(255),
    tag_id VARCHAR(50)
  );
