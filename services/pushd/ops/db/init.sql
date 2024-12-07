CREATE DATABASE glad-pushd
    WITH 
    OWNER = glad_user
    ENCODING = 'UTF8'
    LC_COLLATE = 'en_US.UTF-8'
    LC_CTYPE = 'en_US.UTF-8'
    TEMPLATE = template0
    CONNECTION LIMIT = -1;

-- Connect to the new database (only works in psql)
\c glad-pushd;

-- Create custom ENUM types (add these before the table creation)

-- Create tables
-- CREATE TABLE IF NOT EXISTS tenant (
--     id BIGSERIAL PRIMARY KEY,
--     is_default BOOLEAN UNIQUE,
--     created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
--     updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
-- );

-- PRODUCT entity
-- Note: In Salesforce, this is called as "master". Need to check with PIM experts, but to me
-- Base product (in Salesforce, it is Product) and Product sounds easier to understand.
-- Other possible terminologies are primary product, variants, SKU, etc.
CREATE TABLE IF NOT EXISTS token (
    id SERIAL PRIMARY KEY,
    token VARCHAR(32) NOT NULL UNIQUE,
    -- Note: Do not want to delete tenant if product exists
    -- Tenant can be mapped to organization entity
    tenant_id BIGINT NOT NULL REFERENCES tenant(id),

    -- ext_name is the salesforce's internal name
    ext_name VARCHAR(80) NOT NULL UNIQUE,

    -- User visible name of the product
    title VARCHAR(255) NOT NULL,

    -- Note: Though it appears like a numeric identifier, alpha prefix is present in Salesforce for
    -- this field. Thus, it's marked as a string. Technically, this can be shorter (32 or 16 bytes)
    -- in length.
    ctype VARCHAR(100) NOT NULL,
    
    -- This maps to 'Product' entity in Salesforce via SF's id
    base_product_ext_id VARCHAR(32),

    -- Duration (in days)
    duration_days INTEGER,

    -- Only Public products are made visible on the site. We can filter based on this.
    visibility product_visibility,

    -- maximum attendees
    max_attendees INTEGER,
    format product_format,

    is_auto_approve BOOLEAN DEFAULT FALSE,
    -- is_deleted is an internal field in Salesforce. Hence, need not be synced

    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE INDEX idx_product_ext_id ON product(ext_id);
CREATE INDEX idx_product_tenant_id ON product(tenant_id);
CREATE INDEX idx_product_name ON product(ext_name);
CREATE INDEX idx_product_title ON product(title);
