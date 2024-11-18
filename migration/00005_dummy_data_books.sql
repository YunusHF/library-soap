-- +goose Up
-- +goose StatementBegin
INSERT INTO books (id, title, author, published_date) 
VALUES 
    (1, 'The Great Gatsby', 'F. Scott Fitzgerald', '1925-04-10'),
    (2, 'To Kill a Mockingbird', 'Harper Lee', '1960-07-11'),
    (3, '1984', 'George Orwell', '1949-06-08');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM books WHERE id IN (1,2,3)
-- +goose StatementEnd
