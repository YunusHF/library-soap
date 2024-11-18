-- +goose Up
-- +goose StatementBegin
INSERT INTO customers (id, name)
VALUES
    (121, 'Joni'),
    (122, 'Jono'),
    (123, 'Jona');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM customers WHERE id IN (121,122,123)
-- +goose StatementEnd
