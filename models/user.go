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
	query := `SELECT balance FROM user_data LIMIT 1` // Should only be 1 entry anyways
	err = db.QueryRow(query).Scan(&profit_loss)

	if err == sql.ErrNoRows {
		return 0, nil // No rows found, should never be possible but who knows
	} else if err != nil {
		return 0, err // Other errors
	}

	return profit_loss, nil
}