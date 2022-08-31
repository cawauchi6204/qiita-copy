CREATE TABLE
  IF NOT EXISTS organizations(
    id serial PRIMARY KEY,
    name VARCHAR (50) UNIQUE NOT NULL,
    password VARCHAR (50) NOT NULL,
    email VARCHAR (300) UNIQUE NOT NULL,
    description TEXT NOT NULL,
    hp_url VARCHAR(100) NOT NULL,
    location VARCHAR(50) NOT NULL,
    is_deleted BOOLEAN NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
  );
