-- Create User table --
CREATE TABLE User (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    email TEXT UNIQUE NOT NULL,
    username TEXT UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- Create Category table --
CREATE TABLE Category (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT UNIQUE NOT NULL
);

-- Create Post table --
CREATE TABLE Post (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER NOT NULL,
    title TEXT NOT NULL,
    content TEXT NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    category_id INTEGER,
    FOREIGN KEY (user_id) REFERENCES User(id),
    FOREIGN KEY (category_id) REFERENCES Category(id)
);

-- Create Comment table --
CREATE TABLE Comment (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    post_id INTEGER NOT NULL,
    user_id INTEGER NOT NULL,
    content TEXT NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (post_id) REFERENCES Post(id),
    FOREIGN KEY (user_id) REFERENCES User(id)
);

-- Create LikeDislike table --
CREATE TABLE LikeDislike (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER NOT NULL,
    post_id INTEGER,
    comment_id INTEGER,
    type INTEGER NOT NULL, -- 1 for like, -1 for dislike
    FOREIGN KEY (user_id) REFERENCES User(id),
    FOREIGN KEY (post_id) REFERENCES Post(id),
    FOREIGN KEY (comment_id) REFERENCES Comment(id)
);

DROP TABLE User;
DROP TABLE Category;
DROP TABLE Post;
DROP TABLE Comment;
DROP TABLE LikeDislike;