CREATE TABLE metadata (
    id SERIAL PRIMARY KEY,
    version BIGINT NOT NULL,
    url TEXT NOT NULL,
    last_updated TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    total INTEGER DEFAULT 0,,
    type TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
);

-- Create the trigger function to update last_updated
CREATE OR REPLACE FUNCTION update_last_updated_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.last_updated = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Create a trigger that invokes the above function before any update on the `metadata` table.
CREATE TRIGGER set_last_updated
BEFORE UPDATE ON metadata
FOR EACH ROW
EXECUTE FUNCTION update_last_updated_column();
