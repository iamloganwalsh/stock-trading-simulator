package config

import (
	"database/sql"
	"log"
	_ "github.com/mattn/go-sqlite3"
)

// ConnectDB initializes and returns a database connection
func ConnectDB() (*sql.DB, error) {
    // Open a connection to the SQLite database file
    db, err := sql.Open("sqlite3", "./user_data.db")
    if err != nil {
        log.Printf("Error opening database: %v\n", err)
        return nil, err
    }

    // Verify the connection to the database
    if err := db.Ping(); err != nil {
        log.Printf("Error connecting to the database: %v\n", err)
        return nil, err
    }

    log.Println("Successfully connected to the database.")
    return db, nil
}