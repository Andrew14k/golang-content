package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"
)

var Conn *sql.DB

func New() error {
	const maxAttempts = 5 //5 attempts to run

	//var db *sql.DB //sets up database - pointer to
	var err error

	delay := 2 * time.Second

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", // collects environment variables of db for connection string
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	) // create the db connection string

	for attempts := 1; attempts <= maxAttempts; attempts++ { //tries to make connection to db with a delay if fails before trying again
		log.Printf("Attempt %d: Connecting to database...", attempts)
		Conn, err = sql.Open("mysql", dsn)
		if err == nil {
			err = Conn.Ping()
		}
		if err == nil {
			log.Println("Successfully connected to MySQL!")
			return nil
		}
		log.Printf("Failed to connect: %v", err)
		if attempts < maxAttempts {
			log.Printf("Retrying in %s...\n", delay)
			time.Sleep(delay)
			delay *= 2 // exponential backoff
		}
	}
	return fmt.Errorf("could not connect to database after %d attempts: %w", maxAttempts, err)
}

func Close() {
	Conn.Close()
}
