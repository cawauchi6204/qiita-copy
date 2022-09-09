CREATE TABLE
  IF NOT EXISTS organization_follows(
    id serial PRIMARY KEY,
    follow_user_id VARCHAR(50),
    followed_organization_id VARCHAR(50),
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
  );
