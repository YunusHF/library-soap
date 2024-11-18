package repository

import (
	"database/sql"
	"soap-library/model"
	"time"
)

type LibraryRepo interface {
	FetchBooks() ([]*model.Book, error)
	ValidateCustomer(customerID uint64) (bool, error)
	InsertOrderBooks(product []*model.Product, orderID, customerID uint64) error
	GetStocksByBooksID(booksID uint64) (uint64, error)
}

type libraryRepository struct {
	Db *sql.DB
}

// NewMariaAccountRepository return new maria transaction posting repo
func NewLibraryRepo(
	conn *sql.DB,
) LibraryRepo {
	return &libraryRepository{
		conn,
	}
}

func (lr *libraryRepository) FetchBooks() ([]*model.Book, error) {
	query := "SELECT id, title, author, published_date FROM books"

	rows, err := lr.Db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []*model.Book
	for rows.Next() {
		var book model.Book
		var publishedDate time.Time
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &publishedDate)
		if err != nil {
			return nil, err
		}

		book.PublishedDate = publishedDate.Format("2006-01-02")

		books = append(books, &book)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return books, nil
}

func (lr *libraryRepository) ValidateCustomer(customerID uint64) (bool, error) {
	query := "SELECT id, name FROM customers WHERE id = ?"

	rows := lr.Db.QueryRow(query, customerID)

	var id uint64
	var name string
	err := rows.Scan(&id, &name)
	if err == sql.ErrNoRows {
		return false, nil
	}
	if err != nil {
		return false, err
	}

	return true, nil
}

func (lr *libraryRepository) InsertOrderBooks(product []*model.Product, orderID, customerID uint64) error {
	tx, err := lr.Db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	for _, book := range product {
		// Insert new order
		query := "INSERT INTO transactions (order_id, books_id, quantity, customer_id) VALUES (?, ?, ?, ?)"
		_, err := tx.Exec(query, orderID, book.ID, book.Quantity, customerID)
		if err != nil {
			return err
		}

		// Update stock quantity
		_, err = tx.Exec("UPDATE stocks SET stock = stock - ? WHERE books_id = ?", book.Quantity, book.ID)
		if err != nil {
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (lr *libraryRepository) GetStocksByBooksID(booksID uint64) (uint64, error) {
	query := "SELECT stock FROM stocks WHERE books_id = ?"

	rows := lr.Db.QueryRow(query, booksID)

	var stock uint64
	err := rows.Scan(&stock)
	if err != nil {
		return 0, err
	}

	return stock, nil
}
