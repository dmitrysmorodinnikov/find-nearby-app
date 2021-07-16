CREATE EXTENSION IF NOT EXISTS postgis;
CREATE TABLE locations(vehicle_id INT8 PRIMARY KEY, location GEOMETRY);
CREATE INDEX locations_location_idx ON locations USING GIST (location);