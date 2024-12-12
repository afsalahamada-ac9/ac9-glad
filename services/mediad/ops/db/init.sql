-- Copyright 2024 AboveCloud9.AI Products and Services Private Limited
-- All rights reserved.
-- This code may not be used, copied, modified, or distributed without explicit permission.

CREATE DATABASE glad_mediad
    WITH
    OWNER = glad_user
    ENCODING = 'UTF8'
    LC_COLLATE = 'en_US.UTF-8'
    LC_CTYPE = 'en_US.UTF-8'
    TEMPLATE = template0
    CONNECTION LIMIT = -1;

-- Connect to the new database (only works in psql)
\c glad_mediad;


CREATE TABLE IF NOT EXISTS metadata (
    id SERIAL PRIMARY KEY,
    version INTEGER NOT NULL DEFAULT 1,
    url TEXT NOT NULL,
    total INTEGER NOT NULL,
    type TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    last_updated TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
-- -- Create the trigger function to update last_updated
-- CREATE OR REPLACE FUNCTION update_last_updated_column()
-- RETURNS TRIGGER AS $$
-- BEGIN
--     NEW.last_updated = CURRENT_TIMESTAMP;
--     RETURN NEW;
-- END;
-- $$ LANGUAGE plpgsql;

-- -- Create a trigger that invokes the above function before any update on the `metadata` table.
-- CREATE TRIGGER set_last_updated
-- BEFORE UPDATE ON metadata
-- FOR EACH ROW
-- EXECUTE FUNCTION update_last_updated_column();
