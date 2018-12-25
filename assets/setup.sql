-- Config
CREATE TABLE IF NOT EXISTS config (
    confkey VARCHAR(255),
    confvalue VARCHAR(255),
    PRIMARY KEY (confkey)
);

-- Fill Config
INSERT INTO config (confkey, confvalue)
SELECT 'title', 'Gorum' WHERE NOT EXISTS (
    SELECT confkey FROM config WHERE confkey = 'title'
);

-- Users
CREATE TABLE IF NOT EXISTS users (
    id SERIAL,
    username VARCHAR(255) UNIQUE,
    passwordhash VARCHAR(255),
    registered VARCHAR(255),
    PRIMARY KEY (id)
);

-- Categories
CREATE TABLE IF NOT EXISTS categories (
    id SERIAL,
    categoryname VARCHAR(255) NOT NULL,
    sort INTEGER NOT NULL,
    PRIMARY KEY (id)
);

-- Boards
CREATE TABLE IF NOT EXISTS boards (
    id SERIAL,
    boardname VARCHAR(255) NOT NULL,
    boarddescription VARCHAR(255),
    boardicon VARCHAR(255) DEFAULT 'forum',
    sort INTEGER NOT NULL,
    category INTEGER NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (category) REFERENCES categories(id) ON UPDATE CASCADE ON DELETE CASCADE
);

-- Threads
CREATE TABLE IF NOT EXISTS threads (
    id SERIAL,
    threadname VARCHAR(255) NOT NULL,
    board INTEGER NOT NULL,
    author INTEGER NOT NULL,
    created BIGINT NOT NULL,
    content TEXT NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (board) REFERENCES boards(id) ON UPDATE CASCADE ON DELETE CASCADE,
    FOREIGN KEY (author) REFERENCES users(id) ON UPDATE CASCADE ON DELETE CASCADE
);

-- Posts
CREATE TABLE IF NOT EXISTS posts (
    id SERIAL,
    thread INTEGER NOT NULL,
    author INTEGER NOT NULL,
    created BIGINT NOT NULL,
    content TEXT NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (thread) REFERENCES threads(id) ON UPDATE CASCADE ON DELETE CASCADE,
    FOREIGN KEY (author) REFERENCES users(id) ON UPDATE CASCADE ON DELETE CASCADE
)