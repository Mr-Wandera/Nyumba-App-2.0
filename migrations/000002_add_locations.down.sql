-- Down Migration 000002
ALTER TABLE properties DROP COLUMN IF EXISTS neighborhood_id;
ALTER TABLE properties DROP COLUMN IF EXISTS town_id;
ALTER TABLE properties DROP COLUMN IF EXISTS ward_id;
ALTER TABLE properties DROP COLUMN IF EXISTS sub_county_id;
ALTER TABLE properties DROP COLUMN IF EXISTS county_id;

DROP TABLE IF EXISTS neighborhoods;
DROP TABLE IF EXISTS towns;
DROP TABLE IF EXISTS wards;
DROP TABLE IF EXISTS sub_counties;
DROP TABLE IF EXISTS counties;
