-- migrate:up
CREATE TABLE todo (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    content text NOT NULL,
    checked boolean default 0
);
INSERT INTO todo VALUES ( 1, 'clean room',1)


-- migrate:down
DROP TABLE todo;
