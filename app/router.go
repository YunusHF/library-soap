package app

import (
	"log"
	"net/http"
	"soap-library/pkg/pkgsoap"

	"github.com/gorilla/mux"
)

func (app *App) initRouter() {
	router := mux.NewRouter()

	router.Use(loggingMiddleware)
	router.Use(contentTypeMiddleware)

	router.NotFoundHandler = http.HandlerFunc(func(rw http.ResponseWriter, _ *http.Request) {
		rw.WriteHeader(http.StatusNotFound)
		pkgsoap.SendSOAPError(rw, "Not Found", "not found")
	})

	router.MethodNotAllowedHandler = http.HandlerFunc(func(rw http.ResponseWriter, _ *http.Request) {
		rw.WriteHeader(http.StatusMethodNotAllowed)
		pkgsoap.SendSOAPError(rw, "Method Not Allowed", "method not allowed")
	})

	app.router = router
}

// Middleware for logging
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Incoming %s request to %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

// Middleware for setting content type
func contentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/xml")
		next.ServeHTTP(w, r)
	})
}
