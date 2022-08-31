CREATE TABLE
  IF NOT EXISTS organization_follows(
    id serial PRIMARY KEY,
    follow_user_id INTEGER,
    followed_organization_id INTEGER,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
  );
