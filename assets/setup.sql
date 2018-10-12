-- Users
CREATE TABLE IF NOT EXISTS users (
    id SERIAL,
    username VARCHAR(255),
    passwordhash VARCHAR(255),
    mail VARCHAR(255)
);