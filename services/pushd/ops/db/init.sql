CREATE DATABASE glad_pushd
    WITH
    OWNER = glad_user
    ENCODING = 'UTF8'
    LC_COLLATE = 'en_US.UTF-8'
    LC_CTYPE = 'en_US.UTF-8'
    TEMPLATE = template0
    CONNECTION LIMIT = -1;

-- Connect to the new database (only works in psql)
\c glad_pushd;

-- Create custom ENUM types (add these before the table creation)

CREATE TABLE IF NOT EXISTS device (
    id BIGSERIAL PRIMARY KEY,

    -- Note: Do not want to delete tenant if product exists
    -- Tenant can be mapped to organization entity
    tenant_id BIGINT NOT NULL,
    account_id BIGINT NOT NULL,

    push_token VARCHAR(1024) NOT NULL UNIQUE,
    revoke_id VARCHAR(1024) UNIQUE,
    app_version VARCHAR(16) NOT NULL,

    -- For analytics purposes. Based on the data sent, this may be placed in a different service.
    device_info JSONB,
    platform_info JSONB,

    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
);
CREATE INDEX idx_device_tenant_id ON device(tenant_id);
CREATE INDEX idx_device_account_id ON device(account_id);
CREATE INDEX idx_device_revoke_id ON device(revoke_id);
