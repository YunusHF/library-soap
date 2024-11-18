-- +goose Up
-- +goose StatementBegin
CREATE TABLE stocks (
     `id` SERIAL PRIMARY KEY,
     `books_id` BIGINT UNSIGNED NOT NULL,
     `stock` INT NOT NULL DEFAULT 0
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS stocks;
-- +goose StatementEnd
