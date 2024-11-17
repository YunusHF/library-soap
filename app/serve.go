package app

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func (app *App) Run() {
	terminateChan := app.createTerminateSignal()

	if err := app.Start(); err != nil {
		log.Fatalf("unable to start server, error: %v\n", err)
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
			log.Fatalf("application shutdown with problem, error: %v\n", err)
		}

		terminateChan <- struct{}{}
		close(terminateChan)

		log.Println("application gracefully shutdown")
	}()

	return terminateChan
}
