package delivery

import (
	"context"
	"encoding/xml"
	"log"
	"net/http"
	"soap-library/model"
	"soap-library/pkg/pkgsoap"
	"soap-library/service"

	"github.com/gorilla/mux"
)

type LibraryHandler struct {
	LbService service.LibraryService
}

func NewLibraryHandler(router *mux.Router, libraryService service.LibraryService) {
	handler := &LibraryHandler{
		LbService: libraryService,
	}

	router.HandleFunc("/books", handler.GetBooks).Methods("GET")
	router.HandleFunc("/order", handler.OrderBooks).Methods("POST")
}

func (lb *LibraryHandler) GetBooks(rw http.ResponseWriter, r *http.Request) {
	books, err := lb.LbService.GetBooks(context.Background())
	if err != nil {
		log.Println("[error GetBooks endpoint]:", err)
		pkgsoap.SendSOAPError(rw, "Get Books Error", err.Error())
		return
	}

	response := model.BooksResponse{
		Books: books,
	}

	envelope := pkgsoap.SOAPEnvelope{
		Body: pkgsoap.SOAPBody{
			Content: response,
		},
	}

	pkgsoap.EncodeSOAPResponse(rw, envelope)
}

func (lb *LibraryHandler) OrderBooks(rw http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	var req model.CreateOrderEnvelope

	err := xml.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Println("[error OrderBooks endpoint]: error decode request", err)
		pkgsoap.SendSOAPError(rw, "Invalid Request", err.Error())
		return
	}

	ok, err := lb.LbService.ValidateCustomer(ctx, req.Body.CreateOrderRequest.CustomerID)
	if err != nil {
		pkgsoap.SendSOAPError(rw, "Validate Customer Error", err.Error())
		return
	}
	if !ok {
		pkgsoap.SendSOAPError(rw, "Invalid Customer", err.Error())
		return
	}

	ok, err = lb.LbService.ValidateStocks(ctx, req.Body.CreateOrderRequest.Products)
	if err != nil {
		pkgsoap.SendSOAPError(rw, "Validate Stock Error", err.Error())
		return
	}
	if !ok {
		pkgsoap.SendSOAPError(rw, "Out of Stock", "One or more item is out of stock")
		return
	}

	order, err := lb.LbService.OrderBooks(ctx, model.OrderBooksRequestService{
		CustomerID: req.Body.CreateOrderRequest.CustomerID,
		Products:   req.Body.CreateOrderRequest.Products,
	})
	if err != nil {
		pkgsoap.SendSOAPError(rw, "Order Books Error", err.Error())
		return
	}

	response := model.CreateOrderResponse{
		OrderID: order.OrderID,
		Status:  order.Status,
		Message: order.Message,
	}

	envelope := pkgsoap.SOAPEnvelope{
		Body: pkgsoap.SOAPBody{
			Content: response,
		},
	}

	pkgsoap.EncodeSOAPResponse(rw, envelope)
}
