package app

import (
	"net/http"
	"time"
)

func (app *App) makeHTTPServer() {

	httpServer := &http.Server{
		Addr:         ":8080",
		Handler:      app.router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	app.httpServer = httpServer
}
