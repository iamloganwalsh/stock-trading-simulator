package models

import (
	"database/sql"
)

func InitUser(db *sql.DB, username string) error {

	query := `INSERT INTO user_data (username, balance, profit_loss) VALUES (?, ?, ?)`
	_, err := db.Exec(query, username, 0.0, 0.0)
	return err
}

func GetUsername(db *sql.DB) (string, error) {

	var username string
	query := `SELECT username FROM user_data LIMIT 1` // Should only be 1 entry anyways
	err := db.QueryRow(query).Scan(&username)

	if err == sql.ErrNoRows {
		return "", nil // No rows found, should never be possible but who knows
	} else if err != nil {
		return "", err // Other errors
	}

	return username, nil
}

func GetBalance(db *sql.DB) (float64, error) {

	var balance float64
	query := `SELECT balance FROM user_data LIMIT 1` // Should only be 1 entry anyways
	err := db.QueryRow(query).Scan(&balance)

	if err == sql.ErrNoRows {
		return 0, nil // No rows found, should never be possible but who knows
	} else if err != nil {
		return 0, err // Other errors
	}

	return balance, nil
}

func GetProfitLoss(db *sql.DB) (float64, error) {

	var profit_loss float64
	query := `SELECT profit_loss FROM user_data LIMIT 1` // Should only be 1 entry anyways
	err := db.QueryRow(query).Scan(&profit_loss)

	if err == sql.ErrNoRows {
		return 0, nil // No rows found, should never be possible but who knows
	} else if err != nil {
		return 0, err // Other errors
	}

	return profit_loss, nil
}

type CryptoData struct {
	Code         string  `json:"code"`
	Crypto_count float64 `json:"crypto_count"`
}

func GetCryptoPortfolio(db *sql.DB) ([]CryptoData, error) {
	var crypto_items []CryptoData

	rows, err := db.Query("SELECT code, crypto_count FROM crypto")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var crypto CryptoData
		err = rows.Scan(&crypto.Code, &crypto.Crypto_count)
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
	Stock_count float64 `json:"stock_count"`
}

func GetStockPortfolio(db *sql.DB) ([]StockData, error) {
	var stock_items []StockData

	rows, err := db.Query("SELECT code, stock_count FROM stock")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var stock StockData
		err = rows.Scan(&stock.Code, &stock.Stock_count)
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

type trade_history_data struct {
	Type   string  `json:"type"`
	Code   string  `json:"code"`
	Method string  `json:"method"`
	Cost   float64 `json:"cost"`
	Date   string  `json:"date"`
}

func GetTradeHistory(db *sql.DB) ([]trade_history_data, error) {
	var history_items []trade_history_data

	rows, err := db.Query("SELECT type, code, method, cost, date FROM trade_history")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var history trade_history_data
		err = rows.Scan(&history.Type, &history.Code, &history.Method, &history.Cost, &history.Date)
		if err != nil {
			return nil, err
		}
		history_items = append(history_items, history)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return history_items, nil
}
