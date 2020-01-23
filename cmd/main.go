package main

import (
	"log"
	"os"
	"time"

	"github.com/omerkaya1/didactic-succotash/internal"
)

func main() {
	// Env variables
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPwd := os.Getenv("DB_PWD")
	// TODO: replace with a script that will allow for a more intelligent way of waiting for the DB to be available
	time.Sleep(15 * time.Second)
	// Init Storage
	storage, err := internal.NewStorage(dbName, dbUser, "disable", dbPwd, "postgres", "5432")
	if err != nil {
		log.Fatal(err)
	}
	// Init server
	server := internal.NewServer(storage)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
