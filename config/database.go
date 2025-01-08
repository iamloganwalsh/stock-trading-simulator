package config

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// ConnectDB initializes and returns a database connection
func ConnectDB(testing ...string) (*sql.DB, error) {
	// Open a connection to the SQLite database file

	// If no parameters, we want to use the real DB.
	// If testing contains any values, we want to create a testing database
	db_name := "./user_data.db"
	if len(testing) > 0 {
		db_name = testing[0]
	}

	db, err := sql.Open("sqlite3", db_name)
	if err != nil {
		log.Printf("Error opening database: %v\n", err)
		return nil, err
	}

	// Verify the connection to the database
	if err := db.Ping(); err != nil {
		log.Printf("Error connecting to the database: %v\n", err)
		return nil, err
	}

	//log.Println("Successfully connected to the database.")
	// The above line shows up every time we connect for each function, so I'm commenting it out for now
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
		invested REAL NOT NULL,
		sell_price REAL
	);`

	createStockTableSQL := `
	CREATE TABLE IF NOT EXISTS stock (
		code TEXT NOT NULL,
		amount_held REAL NOT NULL,
		total_bought_cost REAL NOT NULL,
		total_sold_cost REAL NOT NULL,
		total_stock_bought REAL NOT NULL,
		total_stock_sold REAL NOT NULL,
		stock_count REAL NOT NULL
	);`

	createCryptoTableSQL := `
	CREATE TABLE IF NOT EXISTS crypto (
		code TEXT NOT NULL,
		amount_held REAL NOT NULL,
		total_bought_cost REAL NOT NULL,
		total_sold_cost REAL NOT NULL,
		total_crypto_bought REAL NOT NULL,
		total_crypto_sold REAL NOT NULL,
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
