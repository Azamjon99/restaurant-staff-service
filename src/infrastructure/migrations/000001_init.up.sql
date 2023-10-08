-- manager schema
CREATE SCHEMA IF NOT EXISTS manager;

--manager

CREATE TABLE IF NOT EXISTS manager.managers (
	id varchar(36) PRIMARY KEY,
    email VARCHAR(255),
    password_hash VARCHAR(255),
    password_salt VARCHAR(255),
	created_at timestamp,
	updated_at timestamp
);

CREATE INDEX IF NOT EXISTS idx_manager_by_email ON manager.managers(email);

--manager_profiles
CREATE TABLE IF NOT EXISTS manager.manager_profiles (
	manager_id varchar(36) PRIMARY KEY,
	entity_id varchar(36),
	name VARCHAR(255),
    image_url VARCHAR(255),
    phone_number VARCHAR(255),
	email VARCHAR(255),
	created_at timestamp,
	updated_at timestamp
);
CREATE INDEX IF NOT EXISTS idx_manager_by_entity ON manager.manager_profiles(entity_id);
-- manager_restaurant_assignments
CREATE TABLE IF NOT EXISTS manager.manager_restaurant_assignments (
	id serial PRIMARY KEY,
    manager_id VARCHAR(36),
    restaurant_id VARCHAR(36),
	created_at timestamp,
	updated_at timestamp
);

CREATE INDEX IF NOT EXISTS idx_restaurant_by_manager ON manager.manager_restaurant_assignments(manager_id);
CREATE INDEX IF NOT EXISTS idx_manager_by_restaurant ON manager.manager_restaurant_assignments(restaurant_id);
