CREATE TABLE IF NOT EXISTS locations
(
    id         UUID,
    updated_at TIMESTAMP,
    latitude   REAL,
    longitude  REAL,
    PRIMARY KEY (id)
);

ALTER TABLE locations
    OWNER TO pguser;
