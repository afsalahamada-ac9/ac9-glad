-- Copyright 2024 AboveCloud9.AI Products and Services Private Limited
-- All rights reserved.
-- This code may not be used, copied, modified, or distributed without explicit permission.

CREATE DATABASE glad_ldsd
    WITH 
    OWNER = glad_user
    ENCODING = 'UTF8'
    LC_COLLATE = 'en_US.UTF-8'
    LC_CTYPE = 'en_US.UTF-8'
    TEMPLATE = template0
    CONNECTION LIMIT = -1;

-- Connect to the new database (only works in psql)
\c glad_ldsd;


CREATE TABLE live_darshan (
    id BIGINT PRIMARY KEY,
    tenant_id BIGINT,
    date DATE,
    start_time TIMESTAMP,
    meeting_url TEXT,
    created_by BIGINT

    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);