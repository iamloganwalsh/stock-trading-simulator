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

func InitDB(db *sql.DB) error {
	createUserDataTableSQL := `
	CREATE TABLE IF NOT EXISTS user_data (
		username TEXT NOT NULL,
		balance REAL NOT NULL,
		profit_loss REAL NOT NULL
	);`

	createTradeHistoryTableSQL := `
	CREATE TABLE IF NOT EXISTS trade_history (
		type TEXT NOT NULL,
		code TEXT NOT NULL,
		buy_price REAL NOT NULL,
		sell_price REAL
	);`

	createStockTableSQL := `
	CREATE TABLE IF NOT EXISTS stock (
		code TEXT NOT NULL,
		buy_price REAL NOT NULL,
		stock_count REAL NOT NULL
	);`

	createCryptoTableSQL := `
	CREATE TABLE IF NOT EXISTS crypto (
		code TEXT NOT NULL,
		buy_price REAL NOT NULL,
		crypto_count REAL NOT NULL
	);`

	_, err := db.Exec(createUserDataTableSQL)
	if err != nil {
		return err
	}

	_, err = db.Exec(createTradeHistoryTableSQL)
	if err != nil {
		return err
	}

	_, err = db.Exec(createStockTableSQL)
	if err != nil {
		return err
	}

	_, err = db.Exec(createCryptoTableSQL)
	if err != nil {
		return err
	}

	return nil
}
