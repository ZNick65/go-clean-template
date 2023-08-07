CREATE TABLE IF NOT EXISTS clients(
    id serial PRIMARY KEY,
    name text,
    db_name text,
    db_user text,
    db_pass text,
    deleted boolean,
    access_token text
);