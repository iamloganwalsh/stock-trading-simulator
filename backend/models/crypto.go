package models

import (
	"database/sql"
	"fmt"
	"time"
)

func BuyCrypto(db *sql.DB, code string, cost float64, crypto_count float64) error {
	// Start a new transaction
	cost *= crypto_count

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	// Commit the transaction at the end or rollback if there is an error
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	// Check that user can afford
	user_balance, err := GetBalance(db) // From user.go
	if err != nil {
		return err
	}
	if user_balance < cost {
		return fmt.Errorf("insufficient funds")
	}

	var currentCryptoCount float64
	query := `SELECT crypto_count FROM crypto WHERE code = ?`
	err = tx.QueryRow(query, code).Scan(&currentCryptoCount)
	if err != nil {
		if err == sql.ErrNoRows { // Add new entry
			addQuery := `INSERT INTO crypto (code, crypto_count) VALUES (?, ?)`
			_, err = tx.Exec(addQuery, code, crypto_count)
			if err != nil {
				return err
			}
		} else {
			return err
		}
	} else {
		// Update existing entry
		updateQuery := `UPDATE crypto SET crypto_count = ? WHERE code = ?`
		_, err = tx.Exec(updateQuery, crypto_count+currentCryptoCount, code)
		if err != nil {
			return err
		}
	}

	// Update trade history
	tradeCryptoQuery := `INSERT INTO trade_history (type, code, method, cost, date) VALUES (?, ?, ?, ?, ?)`
	_, err = tx.Exec(tradeCryptoQuery, "crypto", code, "buy", cost, time.Now().Format(time.RFC3339))
	if err != nil {
		return err
	}

	// Update user balance
	updateBalanceQuery := `UPDATE user_data SET balance = balance - ? WHERE rowid = 1`
	_, err = tx.Exec(updateBalanceQuery, cost)
	if err != nil {
		return err
	}

	return nil
}

func SellCrypto(db *sql.DB, code string, price float64, sell_quantity float64) error {

	// Start a new transaction
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	var crypto_count float64
	cryptoInfoQuery := `SELECT crypto_count FROM crypto WHERE code = ?`
	err = tx.QueryRow(cryptoInfoQuery, code).Scan(&crypto_count)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("user does not own specified crypto")
		}
		return err
	}

	if sell_quantity > crypto_count {
		return fmt.Errorf("user does not own that much crypto; current amount owned: %.6f", crypto_count)
	}

	// Calculate new crypto count and invested amount
	new_crypto_count := crypto_count - sell_quantity

	if new_crypto_count == 0 {
		// Delete the crypto entry if the new count is zero
		deleteCryptoQuery := `DELETE FROM crypto WHERE code = ?`
		_, err = tx.Exec(deleteCryptoQuery, code)
		if err != nil {
			return err
		}
	} else {
		// Update crypto holdings
		updateCryptoQuery := `UPDATE crypto SET crypto_count = ? WHERE code = ?`
		_, err = tx.Exec(updateCryptoQuery, new_crypto_count, code)
		if err != nil {
			return err
		}
	}

	// Update trade history
	totalCryptoRevenue := price * sell_quantity
	updateTradeQuery := `INSERT INTO trade_history (type, code, method, cost, date) VALUES (?, ?, ?, ?, ?)`
	_, err = tx.Exec(updateTradeQuery, "crypto", code, "sell", totalCryptoRevenue, time.Now().Format(time.RFC3339))

	if err != nil {
		return err
	}

	// Update user balance
	updateBalanceQuery := `UPDATE user_data SET balance = balance + ? WHERE rowid = 1`
	_, err = tx.Exec(updateBalanceQuery, price*sell_quantity)
	if err != nil {
		return err
	}

	return nil
}
