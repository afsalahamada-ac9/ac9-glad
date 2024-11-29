CREATE DATABASE glad
    WITH 
    OWNER = glad_user
    ENCODING = 'UTF8'
    LC_COLLATE = 'en_US.UTF-8'
    LC_CTYPE = 'en_US.UTF-8'
    TEMPLATE = template0
    CONNECTION LIMIT = -1;

-- Connect to the new database (only works in psql)
\c glad;

-- Create custom ENUM types (add these before the table creations)
CREATE TYPE product_visibility AS ENUM ('Public'
    , 'Unlisted'
    );
CREATE TYPE product_format AS ENUM ('In Person'
    , 'Online'
    , 'Destination Retreats'
    );
CREATE TYPE course_status AS ENUM ('draft'
    , 'archived'
    , 'open' 
    , 'expense-submitted'
    , 'expense-declined'
    , 'closed'
    , 'active'
    , 'declined'
    , 'submitted'
    , 'canceled'
    , 'inactive'
    );
CREATE TYPE course_mode AS ENUM ('in-person'
    , 'online');
CREATE TYPE timezone_type AS ENUM ('EST'
    , 'CST'
    , 'MST'
    , 'PST'
    );
CREATE TYPE account_type AS ENUM ('assistant-teacher'
    , 'member'
    , 'organizer'
    , 'student'
    , 'teacher'
    , 'user'
    );
CREATE TYPE center_mode AS ENUM ('in-person'
    , 'online'
    );
CREATE TYPE teaching_eligibility_type AS ENUM ('primary'
    , 'assistant'
    );


-- Create tables
CREATE TABLE IF NOT EXISTS tenant (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    country VARCHAR(128) NOT NULL,
    is_default BOOLEAN UNIQUE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE INDEX idx_tenant_name ON tenant(name);
CREATE INDEX idx_tenant_country ON tenant(country);

-- PRODUCT entity
-- Note: In Salesforce, this is called as "master". Need to check with PIM experts, but to me
-- Base product (in Salesforce, it is Product) and Product sounds easier to understand.
-- Other possible terminologies are primary product, variants, SKU, etc.
CREATE TABLE IF NOT EXISTS product (
    id SERIAL PRIMARY KEY,
    -- Note: ext_id is salesforce id
    ext_id VARCHAR(32) NOT NULL UNIQUE,
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

-- CENTER entity
CREATE TABLE IF NOT EXISTS center (
    id SERIAL PRIMARY KEY,
    -- Note: ext_id is salesforce id
    ext_id VARCHAR(32) NOT NULL UNIQUE,
    -- Note: Do not want to delete tenant if center exists
    tenant_id BIGINT NOT NULL REFERENCES tenant(id),

    -- Note: 'ext_name' is needed in SalesForce. Formula field - Auto generated in SF.
    -- Note: Currently this field is used in typeahead. So, keeping this for now.
    ext_name VARCHAR(80) NOT NULL UNIQUE,

    -- Note: This is the human readable name for the center
    name VARCHAR(255),
    
    -- Note: address format: {"street_1": ..., "street_2": ..., "city": ..., "state": ..., "zip": ..., "country": ...}
    -- When multitenancy is introduced, then country can be removed.
    address JSONB,

    -- Note: geo_location format: {"lat": ..., "long": ...}
    geo_location JSONB,
    -- maximum occupancy
    capacity INTEGER,
    mode center_mode DEFAULT 'in-person',
    webpage VARCHAR(255),
    is_national_center BOOLEAN DEFAULT FALSE,
    is_enabled BOOLEAN,

    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE INDEX idx_center_ext_id ON center(ext_id);
CREATE INDEX idx_center_tenant_id ON center(tenant_id);
CREATE INDEX idx_center_name ON center(name);
CREATE INDEX idx_center_ext_name ON center(ext_name);

CREATE TABLE IF NOT EXISTS center_contact (
    center_id INT NOT NULL REFERENCES center(id) ON DELETE CASCADE,
    name VARCHAR(255),
    phone VARCHAR(32),
    email VARCHAR(80),
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE INDEX idx_center_contact_center_id ON center_contact(center_id);

-- COURSE entity
CREATE TABLE IF NOT EXISTS course (
    id BIGSERIAL PRIMARY KEY,
    -- Note: ext_id is salesforce id
    -- When course is created outside of salesforce, ext_id will be NULL
    ext_id VARCHAR(32) UNIQUE,

    -- Note: Do not want to delete tenant if course exists
    tenant_id BIGINT NOT NULL REFERENCES tenant(id),
    product_id BIGINT NOT NULL REFERENCES product(id),

    name VARCHAR(128) NOT NULL,
    notes TEXT,
    status course_status NOT NULL DEFAULT 'draft',
    max_attendees INTEGER,
    timezone timezone_type,
    -- Note: address format: {"street_1": ..., "street_2": ..., "city": ..., "state": ..., "zip": ..., "country": ...}
    -- When multitenancy is introduced, then country can be removed.
    address JSONB,
    center_id BIGINT NOT NULL REFERENCES center(id) ON DELETE RESTRICT,
    mode course_mode NOT NULL DEFAULT 'in-person',
    num_attendees INTEGER DEFAULT 0,

    -- Note: Details page URL is good enough
    url VARCHAR(512),
    short_url VARCHAR(64),

    -- is_auto_approve does not make sense here. In Salesforce this seems like copied from Master (Product)

    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE INDEX idx_course_tenant_id ON course(tenant_id);
CREATE INDEX idx_course_product_id ON course(product_id);

-- ACCOUNT entity
CREATE TABLE IF NOT EXISTS account (
    id BIGSERIAL PRIMARY KEY,
    -- Note: ext_id is salesforce id
    ext_id VARCHAR(32) NOT NULL UNIQUE,

    -- Note: Do not want to delete tenant if course exists
    tenant_id BIGINT NOT NULL REFERENCES tenant(id),

    -- Authentication id: for US, it's AWS Cognito id
    cognito_id VARCHAR(255),

    -- Note: username is to link the account with the logged in user
    username VARCHAR(128) NOT NULL,
    first_name VARCHAR(40),
    last_name VARCHAR(80),
    phone VARCHAR(32),
    email VARCHAR(80),
    type account_type,

    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
    -- Q: Is it unique? How about kids' account?
    -- UNIQUE(tenant_id, username)
);
CREATE INDEX idx_account_tenant_id ON account(tenant_id);
CREATE INDEX idx_account_username ON account(username);
CREATE INDEX idx_account_type ON account(type);
CREATE INDEX idx_account_cognito_id ON account(cognito_id);
CREATE INDEX idx_account_email ON account(email);

-- COURSE ELIGIBLITY: Courses teacher can teach
CREATE TABLE IF NOT EXISTS teacher_eligibility (
    product_id BIGINT NOT NULL REFERENCES product(id) ON DELETE RESTRICT,
    teacher_id BIGINT NOT NULL REFERENCES account(id) ON DELETE RESTRICT,

    type teaching_eligibility_type,

    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE INDEX idx_teacher_eligibility_product_id ON teacher_eligibility(product_id);
CREATE INDEX idx_teacher_eligibility_teacher_id ON teacher_eligibility(teacher_id);

-- Notes: Max 3 organizers per course
-- Notes: Organizer is not mandatory for a course (Confirm)
CREATE TABLE IF NOT EXISTS course_organizer (
    course_id BIGINT NOT NULL REFERENCES course(id) ON DELETE CASCADE,
    organizer_id BIGINT NOT NULL REFERENCES account(id) ON DELETE RESTRICT,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE INDEX idx_course_organizer_course_id ON course_organizer(course_id);
CREATE INDEX idx_course_organizer_organizer_id ON course_organizer(organizer_id);

-- Notes: Max 1 contact per course
-- Notes: Contact is not mandatory for a course
CREATE TABLE IF NOT EXISTS course_contact (
    course_id BIGINT NOT NULL REFERENCES course(id) ON DELETE CASCADE,
    contact_id BIGINT NOT NULL REFERENCES account(id) ON DELETE RESTRICT,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE INDEX idx_course_contact_course_id ON course_contact(course_id);
CREATE INDEX idx_course_contact_contact_id ON course_contact(contact_id);

-- Notes: Max 3 teachers per course
-- Notes: As per SF data model, even teachers are optional for a course
CREATE TABLE IF NOT EXISTS course_teacher (
    course_id BIGINT NOT NULL REFERENCES course(id) ON DELETE CASCADE,
    teacher_id BIGINT NOT NULL REFERENCES account(id) ON DELETE RESTRICT,
    is_primary BOOLEAN DEFAULT FALSE,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE INDEX idx_course_teacher_course_id ON course_teacher(course_id);
CREATE INDEX idx_course_teacher_teacher_id ON course_teacher(teacher_id);

-- Note: Tenant is not required for course_timing
CREATE TABLE IF NOT EXISTS course_timing (
    id BIGSERIAL PRIMARY KEY,
    course_id BIGINT NOT NULL REFERENCES course(id) ON DELETE CASCADE,
    -- Note: ext_id is salesforce id
    ext_id VARCHAR(32) UNIQUE,
    -- Note: 'name' is needed in SalesForce. It's a formula field and hence not used here.
    -- Note: 'timezone' is needed in SalesForce. Set the value from 'course' table
    course_date DATE,
    start_time TIME,
    end_time TIME,

    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE INDEX idx_course_timings_course_id ON course_timing(course_id);

-- Note: Tenant is not required for course_notify
-- Notes: Notify is not mandatory for a course
CREATE TABLE IF NOT EXISTS course_notify (
    course_id BIGINT NOT NULL REFERENCES course(id) ON DELETE CASCADE,
    notify_id BIGINT NOT NULL REFERENCES account(id) ON DELETE RESTRICT,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE INDEX idx_course_notify_course_id ON course_notify(course_id);
