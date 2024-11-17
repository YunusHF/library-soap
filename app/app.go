package app

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"net/http"
	"soap-library/delivery"
	"soap-library/pkg/pkgcmd"
	"soap-library/pkg/pkgtime"
	"soap-library/repository"
	"soap-library/service"

	"github.com/gorilla/mux"
)

var _ pkgcmd.Runnable = (*App)(nil)

type App struct {
	database   *sql.DB
	router     *mux.Router
	httpServer *http.Server
	clock      pkgtime.Time
	closersFn  []func(context.Context) error
}

func (app *App) Start() error {
	log.Println("server starting...")

	go func() {
		if app.httpServer == nil {
			return
		}

		log.Println("http server listen on", app.httpServer.Addr)
		if err := app.httpServer.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("fail to start rest server, err: %v\n", err)
		}
	}()

	return nil
}

func (app *App) Stop(ctx context.Context) error {
	for _, closer := range app.closersFn {
		if err := closer(ctx); err != nil {
			return err
		}
	}

	return nil
}

func NewServer() (*App, error) {
	app := &App{
		clock: pkgtime.NewTime(),
	}

	return app.spinUp()
}

func (app *App) spinUp() (*App, error) {

	if err := app.initDB(); err != nil {
		return nil, err
	}

	app.initRouter()
	app.makeHTTPServer()
	app.setUpClosers()

	// init repository
	libraryRepo := repository.NewLibraryRepo(app.database)

	//init service
	libraryService := service.NewLibraryService(libraryRepo)

	// init handlet
	delivery.NewLibraryHandler(app.router, libraryService)

	return app, nil
}

func (app *App) setUpClosers() {
	app.closersFn = append([]func(context.Context) error{
		func(ctx context.Context) error {
			log.Println("stopping http")
			return app.httpServer.Shutdown(ctx)
		},
	}, app.closersFn...)
}
