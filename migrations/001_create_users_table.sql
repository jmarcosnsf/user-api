-- Write your migrate up statements here
CREATE TABLE users (
  id UUID PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  email VARCHAR(255) NOT NULL UNIQUE, 
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);
---- create above / drop below ----
DROP TABLE users;
-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
