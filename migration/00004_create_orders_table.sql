-- +goose Up
-- +goose StatementBegin
CREATE TABLE transactions (
    `id` SERIAL PRIMARY KEY,
    `order_id` BIGINT UNSIGNED NOT NULL,
    `books_id` BIGINT UNSIGNED NOT NULL,
    `customer_id` BIGINT UNSIGNED NOT NULL,
    `quantity` INT NOT NULL DEFAULT 0
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS orders
-- +goose StatementEnd