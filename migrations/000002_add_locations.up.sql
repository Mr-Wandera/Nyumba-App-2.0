-- Migration 000002: Comprehensive Kenya Location Hierarchy
-- DBMS: PostgreSQL

CREATE TABLE IF NOT EXISTS counties (
    id SERIAL PRIMARY KEY,
    code INT UNIQUE NOT NULL,
    name VARCHAR(100) UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS sub_counties (
    id SERIAL PRIMARY KEY,
    county_id INT NOT NULL REFERENCES counties(id) ON DELETE CASCADE,
    name VARCHAR(100) NOT NULL,
    UNIQUE(county_id, name)
);

CREATE TABLE IF NOT EXISTS wards (
    id SERIAL PRIMARY KEY,
    sub_county_id INT NOT NULL REFERENCES sub_counties(id) ON DELETE CASCADE,
    name VARCHAR(100) NOT NULL,
    UNIQUE(sub_county_id, name)
);

CREATE TABLE IF NOT EXISTS towns (
    id SERIAL PRIMARY KEY,
    county_id INT NOT NULL REFERENCES counties(id) ON DELETE CASCADE,
    name VARCHAR(100) NOT NULL,
    UNIQUE(county_id, name)
);

CREATE TABLE IF NOT EXISTS neighborhoods (
    id SERIAL PRIMARY KEY,
    sub_county_id INT REFERENCES sub_counties(id) ON DELETE CASCADE,
    town_id INT REFERENCES towns(id) ON DELETE CASCADE,
    name VARCHAR(100) NOT NULL
);

-- Add location hierarchy columns to properties table
ALTER TABLE properties ADD COLUMN IF NOT EXISTS county_id INT REFERENCES counties(id) ON DELETE SET NULL;
ALTER TABLE properties ADD COLUMN IF NOT EXISTS sub_county_id INT REFERENCES sub_counties(id) ON DELETE SET NULL;
ALTER TABLE properties ADD COLUMN IF NOT EXISTS ward_id INT REFERENCES wards(id) ON DELETE SET NULL;
ALTER TABLE properties ADD COLUMN IF NOT EXISTS town_id INT REFERENCES towns(id) ON DELETE SET NULL;
ALTER TABLE properties ADD COLUMN IF NOT EXISTS neighborhood_id INT REFERENCES neighborhoods(id) ON DELETE SET NULL;

CREATE INDEX IF NOT EXISTS idx_properties_county ON properties(county_id);
CREATE INDEX IF NOT EXISTS idx_properties_sub_county ON properties(sub_county_id);
CREATE INDEX IF NOT EXISTS idx_properties_town ON properties(town_id);
CREATE INDEX IF NOT EXISTS idx_sub_counties_county ON sub_counties(county_id);
CREATE INDEX IF NOT EXISTS idx_wards_sub_county ON wards(sub_county_id);
CREATE INDEX IF NOT EXISTS idx_towns_county ON towns(county_id);
