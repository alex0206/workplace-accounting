CREATE TABLE IF NOT EXISTS workplace
(
    id         serial PRIMARY KEY,
    info       jsonb NOT NULL,
    updated_at timestamp default current_timestamp,
    created_at timestamp default current_timestamp
);
