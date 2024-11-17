package main

import (
	"log"
	"soap-library/app"
)

func main() {
	server, err := app.NewServer()
	if err != nil {
		log.Fatalf("Unable to initiate the server, error: %v\n", err)
	}

	server.Run()
}
