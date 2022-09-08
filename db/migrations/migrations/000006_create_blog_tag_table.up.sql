CREATE TABLE
  IF NOT EXISTS blog_tags(
    id serial PRIMARY KEY,
    post_id VARCHAR(50),
    tag_id VARCHAR(50)
  );
