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
