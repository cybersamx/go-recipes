CREATE TABLE IF NOT EXISTS bus_stops
(
    id         SERIAL PRIMARY KEY ,
    updated_at TIMESTAMP,
    number     VARCHAR(16),
    latitude   REAL,
    longitude  REAL,
    siteats    VARCHAR(64),
    city_site  VARCHAR(48)
);

ALTER TABLE bus_stops
    OWNER TO pguser;
