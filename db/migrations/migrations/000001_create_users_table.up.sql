CREATE TABLE
  IF NOT EXISTS users(
    id serial PRIMARY KEY,
    name VARCHAR (50) NOT NULL,
    password VARCHAR (300) NOT NULL,
    email VARCHAR (300) UNIQUE NOT NULL,
    nickname VARCHAR(50) NOT NULL,
    description TEXT NOT NULL,
    hp_url VARCHAR(100) NOT NULL,
    location VARCHAR(50) NOT NULL,
    github_account VARCHAR(100) NOT NULL,
    organization_id INTEGER,
    is_deleted BOOLEAN NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
  );
