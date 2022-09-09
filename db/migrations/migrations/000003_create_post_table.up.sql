CREATE TABLE
  IF NOT EXISTS posts(
    id VARCHAR(255),
    title VARCHAR (255) NOT NULL,
    body TEXT NOT NULL,
    posted_by VARCHAR(50) NOT NULL,
    organization_id VARCHAR(50),
    is_draft BOOLEAN NOT NULL,
    is_deleted BOOLEAN NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
  );
