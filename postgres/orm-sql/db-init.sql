CREATE TABLE IF NOT EXISTS users
(
    id   SERIAL PRIMARY KEY,
    name VARCHAR(64),
    age  INTEGER
);

ALTER TABLE users
    OWNER TO pguser;

CREATE TABLE IF NOT EXISTS restaurants
(
    id         SERIAL PRIMARY KEY ,
    user_id    INTEGER REFERENCES users,
    visited_at TIMESTAMP,
    name       VARCHAR(64),
    num_seats  INTEGER,
    latitude   REAL,
    longitude  REAL
);

ALTER TABLE restaurants
    OWNER TO pguser;
