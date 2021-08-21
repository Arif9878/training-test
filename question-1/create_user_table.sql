CREATE TABLE IF NOT EXISTS USERS (
    id SERIAL,
    username CHAR varying(15),
    parent INT
);