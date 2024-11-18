package service

import (
	"context"
	"database/sql"
	"log"
	"soap-library/model"
	"soap-library/pkg/pkguid"
	"soap-library/repository"
)

type LibraryService interface {
	GetBooks(ctx context.Context) ([]*model.Book, error)
	OrderBooks(ctx context.Context, req model.OrderBooksRequestService) (*model.OrderBooksResponseService, error)
	ValidateCustomer(ctx context.Context, customerID uint64) (bool, error)
	ValidateStocks(ctx context.Context, product []*model.Product) (bool, error)
}

type Library struct {
	repo repository.LibraryRepo
	sf   pkguid.Snowflake
}

func NewLibraryService(
	repo repository.LibraryRepo,
	sf pkguid.Snowflake,
) LibraryService {
	return &Library{
		repo: repo,
		sf:   sf,
	}
}

func (lb *Library) GetBooks(ctx context.Context) ([]*model.Book, error) {
	books, err := lb.repo.FetchBooks()
	if err != nil {
		log.Println("[error GetBooks service]:", err)
		return nil, err
	}

	return books, nil
}

func (lb *Library) OrderBooks(ctx context.Context, req model.OrderBooksRequestService) (*model.OrderBooksResponseService, error) {

	orderID := lb.sf.GenerateSfID().Uint64()
	err := lb.repo.InsertOrderBooks(req.Products, orderID, req.CustomerID)
	if err != nil {
		log.Println("[error GetBooks service]:", err)
		return &model.OrderBooksResponseService{
			OrderID: 0,
			Status:  "Failed",
			Message: "Failed Create Order",
		}, err
	}

	return &model.OrderBooksResponseService{
		OrderID: orderID,
		Status:  "Success",
		Message: "Order Created Successfully",
	}, nil
}

func (lb *Library) ValidateCustomer(ctx context.Context, customerID uint64) (bool, error) {
	ok, err := lb.repo.ValidateCustomer(customerID)
	if err != nil {
		log.Println("[error GetBooks service]:", err)
		return false, err
	}
	return ok, nil
}

func (lb *Library) ValidateStocks(ctx context.Context, product []*model.Product) (bool, error) {
	for _, book := range product {
		stock, err := lb.repo.GetStocksByBooksID(book.ID)
		if err != nil {
			if err == sql.ErrNoRows {
				log.Println("[error ValidateStocks service]:", err)
				return false, nil
			}
			log.Println("[error ValidateStocks service]:", err)
			return false, err
		}
		if stock < book.Quantity {
			log.Println("[error ValidateStocks service]:", "one or more items is out of stock")
			return false, nil
		}
	}

	return true, nil
}
