CREATE TABLE IF NOT EXISTS "items" (
    id serial PRIMARY KEY,
    created_at      timestamp with time zone default now() not null,
    updated_at      timestamp with time zone,
    name varchar(255) not null,
    price integer not null
);