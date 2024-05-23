CREATE TABLE comments
(
    id        SERIAL PRIMARY KEY,
    user_id   INTEGER                          NOT NULL,
    post_id   INTEGER REFERENCES posts (id)    NOT NULL,
    parent_id INTEGER REFERENCES comments (id) NULL,
    body      VARCHAR(2000)                    NOT NULL
);