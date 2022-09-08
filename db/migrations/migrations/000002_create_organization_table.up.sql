CREATE TABLE
  IF NOT EXISTS organizations(
    id VARCHAR (50) NOT NULL,
    name VARCHAR (50) NOT NULL,
    password VARCHAR (300) NOT NULL,
    email VARCHAR (300) UNIQUE NOT NULL,
    description TEXT NOT NULL,
    hp_url VARCHAR(100) NOT NULL,
    location VARCHAR(50) NOT NULL,
    is_deleted BOOLEAN NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
  );
