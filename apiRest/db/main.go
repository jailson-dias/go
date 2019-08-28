package db

import (
	"fmt"
	"log"
	"os"

	"github.com/go-pg/pg"
)

func Connect() *pg.DB {
	fmt.Println("Tsasds connect")
	opts := &pg.Options{
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Addr:     os.Getenv("POSTGRES_DB"),
	}

	var db *pg.DB = pg.Connect(opts)

	if db == nil {
		log.Println("Failed to connect to database.")
		os.Exit(1)
	}
	log.Println("Connection to database successful.")

	return db
}

func Close(db *pg.DB) {
	closeErr := db.Close()

	if closeErr != nil {
		log.Printf("Error while closing the connection, reason: %v\n", closeErr)
		os.Exit(1)
	}
	log.Println("Connection closed successfully.")
}
