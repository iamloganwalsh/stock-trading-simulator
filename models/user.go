package models

import (
	"database/sql"

	"github.com/iamloganwalsh/stock-trading-simulator/config"
)

func InitUser(username string) error {
	db, err := config.ConnectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	query := `INSERT INTO user_data (username, balance, profit_loss) VALUES (?, ?, ?)`
	_, err = db.Exec(query, username, 0.0, 0.0)
	return err
}

func GetUsername() (string, error) {
	db, err := config.ConnectDB()
	if err != nil {
		return "", err
	}
	defer db.Close()

	var username string
	query := `SELECT username FROM user_data LIMIT 1` // Should only be 1 entry anyways
	err = db.QueryRow(query).Scan(&username)

	if err == sql.ErrNoRows {
		return "", nil // No rows found, should never be possible but who knows
	} else if err != nil {
		return "", err // Other errors
	}

	return username, nil
}

func GetBalance() (float64, error) {
	db, err := config.ConnectDB()
	if err != nil {
		return 0, err
	}
	defer db.Close()

	var balance float64
	query := `SELECT balance FROM user_data LIMIT 1` // Should only be 1 entry anyways
	err = db.QueryRow(query).Scan(&balance)

	if err == sql.ErrNoRows {
		return 0, nil // No rows found, should never be possible but who knows
	} else if err != nil {
		return 0, err // Other errors
	}

	return balance, nil
}

func GetProfitLoss() (float64, error) {
	db, err := config.ConnectDB()
	if err != nil {
		return 0, err
	}
	defer db.Close()

	var profit_loss float64
	query := `SELECT profit_loss FROM user_data LIMIT 1` // Should only be 1 entry anyways
	err = db.QueryRow(query).Scan(&profit_loss)

	if err == sql.ErrNoRows {
		return 0, nil // No rows found, should never be possible but who knows
	} else if err != nil {
		return 0, err // Other errors
	}

	return profit_loss, nil
}

type CryptoData struct {
	Code         string  `json:"code"`
	Invested     float64 `json:"invested"`
	Crypto_count float64 `json:"crypto_count"`
}

func GetCryptoPortfolio() ([]CryptoData, error) {
	var crypto_items []CryptoData

	db, err := config.ConnectDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT code, invested, crypto_count FROM crypto")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var crypto CryptoData
		err = rows.Scan(&crypto.Code, &crypto.Invested, &crypto.Crypto_count)
		if err != nil {
			return nil, err
		}
		crypto_items = append(crypto_items, crypto)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return crypto_items, nil
}

type StockData struct {
	Code        string  `json:"code"`
	Invested    float64 `json:"invested"`
	Stock_count float64 `json:"stock_count"`
}

func GetStockPortfolio() ([]StockData, error) {
	var stock_items []StockData

	db, err := config.ConnectDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT code, invested, stock_count FROM stock")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var stock StockData
		err = rows.Scan(&stock.Code, &stock.Invested, &stock.Stock_count)
		if err != nil {
			return nil, err
		}
		stock_items = append(stock_items, stock)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return stock_items, nil
}
