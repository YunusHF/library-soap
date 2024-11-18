package app

import (
	"context"
	"errors"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/hashicorp/go-multierror"
)

func (app *App) Run() {
	terminateChan := app.createTerminateSignal()

	if err := app.Start(); app.err != nil || err != nil {
		var mErr *multierror.Error
		if errors.As(app.err, &mErr) {
			log.Fatalf("error: %+v", mErr.ErrorOrNil())
		}
	}

	<-terminateChan
}

func (app *App) createTerminateSignal() <-chan struct{} {
	var terminateChan = make(chan struct{})

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

		sigReceived := <-sigint

		log.Println("received signal:", sigReceived)

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := app.Stop(ctx); err != nil {
			app.err = multierror.Append(app.err, err)
		}

		terminateChan <- struct{}{}
		close(terminateChan)

		log.Println("application gracefully shutdown")
	}()

	return terminateChan
}
