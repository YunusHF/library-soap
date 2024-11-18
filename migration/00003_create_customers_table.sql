-- +goose Up
-- +goose StatementBegin
CREATE TABLE customers (
    `id` SERIAL PRIMARY KEY,
    `name` VARCHAR(255) NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS customers
-- +goose StatementEnd
