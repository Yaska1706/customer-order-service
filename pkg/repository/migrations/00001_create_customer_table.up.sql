CREATE TABLE IF EXISTS "customers" (
    id serial PRIMARY KEY
    created_at      timestamp with time zone default now() not null,
    updated_at      timestamp with time zone,
    firstname varchar(255) not null,
    lastname varchar(255) not null,
    username varchar(255) not null,
    password varchar(255) not null,
    phonenumber varchar(255) unique not null,
)