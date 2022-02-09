CREATE TABLE IF NOT EXISTS "orders" (
    id serial PRIMARY KEY,
    created_at      timestamp with time zone default now() not null,
    updated_at      timestamp with time zone,
    amount integer not null,
    customer_id integer not null,
    FOREIGN KEY (customer_id) REFERENCES "customers" (id),
    item_id integer not null,
    FOREIGN KEY (item_id) REFERENCES "items" (id)
);