CREATE TABLE IF NOT EXISTS users
(
    id         INT PRIMARY KEY,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    age        INT,
    email      VARCHAR(128) UNIQUE,
    username   VARCHAR(32) UNIQUE,
    first_name VARCHAR(64),
    last_name  VARCHAR(64)
);

ALTER TABLE users
    OWNER TO pguser;
