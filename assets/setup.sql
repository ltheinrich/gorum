-- Users
CREATE TABLE IF NOT EXISTS users (
    id SERIAL NOT NULL,
    username VARCHAR(255) UNIQUE NOT NULL,
    passwordhash VARCHAR(255) NOT NULL,
    registered VARCHAR(255) NOT NULL,
    PRIMARY KEY (id)
);

-- Tokens
CREATE TABLE IF NOT EXISTS tokens (
    token VARCHAR(255) UNIQUE NOT NULL,
    holder INTEGER NOT NULL,
    created BIGINT NOT NULL,
    PRIMARY KEY (token),
    FOREIGN KEY (holder) REFERENCES users(id) ON UPDATE CASCADE ON DELETE CASCADE
);

-- User Data
CREATE TABLE IF NOT EXISTS userdata (
    holder INTEGER NOT NULL,
    dataname VARCHAR(255) NOT NULL,
    datavalue TEXT NOT NULL,
    UNIQUE(holder, dataname)
);

-- Categories
CREATE TABLE IF NOT EXISTS categories (
    id SERIAL NOT NULL,
    categoryname VARCHAR(255) NOT NULL,
    sort INTEGER NOT NULL,
    PRIMARY KEY (id)
);

-- Boards
CREATE TABLE IF NOT EXISTS boards (
    id SERIAL NOT NULL,
    boardname VARCHAR(255) NOT NULL,
    boarddescription VARCHAR(255),
    boardicon VARCHAR(255) DEFAULT 'forum' NOT NULL,
    sort INTEGER NOT NULL,
    category INTEGER NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (category) REFERENCES categories(id) ON UPDATE CASCADE ON DELETE CASCADE
);

-- Threads
CREATE TABLE IF NOT EXISTS threads (
    id SERIAL NOT NULL,
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
    id SERIAL NOT NULL,
    thread INTEGER NOT NULL,
    author INTEGER NOT NULL,
    created BIGINT NOT NULL,
    content TEXT NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (thread) REFERENCES threads(id) ON UPDATE CASCADE ON DELETE CASCADE,
    FOREIGN KEY (author) REFERENCES users(id) ON UPDATE CASCADE ON DELETE CASCADE
)
