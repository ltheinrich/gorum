-- Config
CREATE TABLE IF NOT EXISTS config (
    confkey VARCHAR(255),
    confvalue VARCHAR(255)
);

-- Fill Config
INSERT INTO config
    (confkey, confvalue)
SELECT 'title', 'Gorum'
WHERE
    NOT EXISTS (
        SELECT confkey FROM config WHERE confkey = 'title'
    );

-- Users
CREATE TABLE IF NOT EXISTS users (
    id SERIAL,
    username VARCHAR(255),
    passwordhash VARCHAR(255),
    mail VARCHAR(255)
);