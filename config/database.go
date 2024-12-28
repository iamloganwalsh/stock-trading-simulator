package config

import (
    "database/sql"
    _ "github.com/lib/pq" // PostgreSQL driver
    "log"
    "fmt"
)

var db *sql.DB

// ConnectToDB sets up the connection to the PostgreSQL database
func ConnectToDB() {
	// Connection string
	connStr := "user=postgres password=12345 dbname=stock_trading_simulator_db sslmode=disable host=localhost"
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}

	// Test the connection
	err = db.Ping()
	if err != nil {
		log.Fatal("Unable to reach the database: ", err)
	}

	fmt.Println("Successfully connected to the database!")
}

// GetDB returns the database connection
func GetDB() *sql.DB {
	return db
}

