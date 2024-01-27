-- +goose Up
BEGIN;

CREATE TABLE customers (
    user_id SERIAL PRIMARY KEY,
    email VARCHAR(50) NOT NULL,
    password VARCHAR(50) NOT NULL,
    name VARCHAR(100) NOT NULL
);

INSERT INTO customers (email, password, name)
VALUES
    ('ivan', '12345', 'Ivan Ivanovich'),
    ('petr', '54321', 'Petr Petrovich');

COMMIT;

-- +goose Down
DROP TABLE company;
