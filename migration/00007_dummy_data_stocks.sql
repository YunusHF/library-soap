-- +goose Up
-- +goose StatementBegin
INSERT INTO stocks (id, books_id, stock)
VALUES
    (1, 1, 10),
    (2, 2, 5),
    (3, 3, 15);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM stocks WHERE id IN (1,2,3)
-- +goose StatementEnd
