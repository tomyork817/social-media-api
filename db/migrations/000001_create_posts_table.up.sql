CREATE TABLE posts
(
    id      SERIAL PRIMARY KEY,
    user_id INTEGER       NOT NULL,
    body    VARCHAR(3000) NOT NULL,
    is_open BOOLEAN       NOT NULL
);