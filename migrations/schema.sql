CREATE TABLE IF NOT EXISTS "schema_migrations" (version varchar(255) primary key);
CREATE TABLE todo (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    content text NOT NULL,
    checked boolean default 0
);
CREATE TABLE sqlite_sequence(name,seq);
-- Dbmate schema migrations
INSERT INTO "schema_migrations" (version) VALUES
  ('20230308195639');
