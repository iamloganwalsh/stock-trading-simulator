package models

import (
	"database/sql"

	"github.com/iamloganwalsh/stock-trading-simulator/config"
)

// If its a disaster consider using a transaction
func BuyCrypto(code string, cost float64, cryptoCount float64) error {
	db, err := config.ConnectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	var invested float64
	var oldCount float64
	query := `SELECT invested, crypto_count FROM crypto WHERE code = ?`
	err = db.QueryRow(query, code).Scan(&invested, &oldCount)
	if err != nil {
		if err == sql.ErrNoRows { // Add new entry
			addQuery := `INSERT INTO crypto (code, invested, crypto_count) VALUES (?, ?, ?)`
			_, err = db.Exec(addQuery, code, cost, cryptoCount)
			return err
		}
		return err
	}

	// Update existing entry
	updateQuery := `UPDATE crypto SET invested = ?, crypto_count = ? WHERE code = ?`
	_, err = db.Exec(updateQuery, invested+cost, cryptoCount+oldCount, code)
	return err
}
