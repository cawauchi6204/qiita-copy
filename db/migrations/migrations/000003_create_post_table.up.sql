CREATE TABLE
  IF NOT EXISTS posts(
    id serial PRIMARY KEY,
    title VARCHAR (50) UNIQUE NOT NULL,
    body TEXT NOT NULL,
    posted_by INTEGER NOT NULL,
    organization_id INTEGER,
    is_draft BOOLEAN NOT NULL,
    is_deleted BOOLEAN NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
  );
