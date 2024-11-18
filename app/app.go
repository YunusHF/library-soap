package app

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"net/http"
	"soap-library/delivery"
	"soap-library/pkg/pkgcmd"
	"soap-library/pkg/pkguid"
	"soap-library/repository"
	"soap-library/service"

	"github.com/gorilla/mux"
	"github.com/hashicorp/go-multierror"
)

var _ pkgcmd.Runnable = (*App)(nil)

type App struct {
	database   *sql.DB
	router     *mux.Router
	httpServer *http.Server
	snowflake  pkguid.Snowflake
	err        error
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
			app.err = multierror.Append(app.err, err)
			log.Fatalf("fail to start rest server, err: %v\n", err)
		}
	}()

	return nil
}

func (app *App) Stop(ctx context.Context) error {
	for _, closer := range app.closersFn {
		if err := closer(ctx); err != nil {
			app.err = multierror.Append(app.err, err)
			return err
		}
	}

	return nil
}

func NewServer() (*App, error) {
	app := &App{
		err: nil,
	}

	return app.spinUp()
}

func (app *App) spinUp() (*App, error) {

	if err := app.initDB(); err != nil {
		app.err = multierror.Append(app.err, err)
		return nil, err
	}

	app.initRouter()
	app.makeHTTPServer()
	app.initSnowflakeGen()
	app.setUpClosers()

	// init repository
	libraryRepo := repository.NewLibraryRepo(app.database)

	//init service
	libraryService := service.NewLibraryService(libraryRepo, app.snowflake)

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
