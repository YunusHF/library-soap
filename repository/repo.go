package repository

import (
	"database/sql"
	"soap-library/model"
	"time"
)

type LibraryRepo interface {
	FetchBooks() ([]*model.Book, error)
}

type libraryRepository struct {
	Db *sql.DB
}

// NewMariaAccountRepository return new maria transaction posting repo
func NewLibraryRepo(Conn *sql.DB) LibraryRepo {
	return &libraryRepository{Conn}
}

func (lr *libraryRepository) FetchBooks() ([]*model.Book, error) {
	rows, err := lr.Db.Query("SELECT id, title, author, published_date FROM books")
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
