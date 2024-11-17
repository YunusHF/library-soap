package delivery

import (
	"context"
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
}

func (lb *LibraryHandler) GetBooks(rw http.ResponseWriter, r *http.Request) {
	books, err := lb.LbService.GetBooks(context.Background())
	if err != nil {
		pkgsoap.SendSOAPError(rw, "Database error", err.Error())
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
