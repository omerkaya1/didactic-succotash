package main

import (
	"log"

	"github.com/omerkaya1/didactic-succotash/internal"
)

func main() {
	// Init Storage
	storage, err := internal.NewStorage("")
	if err != nil {
		log.Fatal(err)
	}
	// Init server
	server := internal.NewServer(storage)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
