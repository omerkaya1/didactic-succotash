package main

import "log"

func main() {
	// Init Storage
	storage, err := NewStorage("")
	if err != nil {
		log.Fatal(err)
	}
	// Init server
	server := NewServer(storage)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
