CREATE TABLE
  IF NOT EXISTS blog_tags(
    id serial PRIMARY KEY,
    post_id INTEGER,
    tag_id INTEGER
  );
