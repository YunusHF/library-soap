package service

import (
	"context"
	"soap-library/model"
	"soap-library/repository"
)

type LibraryService interface {
	GetBooks(ctx context.Context) ([]*model.Book, error)
}

type Library struct {
	repo repository.LibraryRepo
}

func NewLibraryService(repo repository.LibraryRepo) LibraryService {
	return &Library{
		repo: repo,
	}
}

func (lb *Library) GetBooks(ctx context.Context) ([]*model.Book, error) {
	books, err := lb.repo.FetchBooks()
    if err != nil {
        return nil, err
    }

	return books, nil
}