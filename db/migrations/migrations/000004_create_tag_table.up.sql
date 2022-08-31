CREATE TABLE
  IF NOT EXISTS tags(
    id serial PRIMARY KEY,
    name VARCHAR (50) UNIQUE NOT NULL,
    posted_by INTEGER NOT NULL,
    is_deleted BOOLEAN NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
  );
