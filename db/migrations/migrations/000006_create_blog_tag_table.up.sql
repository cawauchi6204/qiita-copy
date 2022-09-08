CREATE TABLE
  IF NOT EXISTS blog_tags(
    id serial PRIMARY KEY,
    post_id VARCHAR(255),
    tag_id VARCHAR(50)
  );
