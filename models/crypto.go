package models

import (
	"database/sql"
	"fmt"

	"github.com/iamloganwalsh/stock-trading-simulator/config"
)

func BuyCrypto(code string, cost float64, crypto_count float64) error {
	// Start a new transaction
	cost *= crypto_count
	db, err := config.ConnectDB()
	if err != nil {
		return err
	}
	defer db.Close()

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

	var invested float64
	var old_count float64
	query := `SELECT invested, crypto_count FROM crypto WHERE code = ?`
	err = tx.QueryRow(query, code).Scan(&invested, &old_count)
	if err != nil {
		if err == sql.ErrNoRows { // Add new entry
			addQuery := `INSERT INTO crypto (code, invested, crypto_count) VALUES (?, ?, ?)`
			_, err = tx.Exec(addQuery, code, cost, crypto_count)
			if err != nil {
				return err
			}
		} else {
			return err
		}
	} else {
		// Update existing entry
		updateQuery := `UPDATE crypto SET invested = ?, crypto_count = ? WHERE code = ?`
		_, err = tx.Exec(updateQuery, invested+cost, crypto_count+old_count, code)
		if err != nil {
			return err
		}
	}

	// Double check that user can afford
	user_balance, err := GetBalance() // From user.go
	if err != nil {
		return err
	}
	if user_balance < cost {
		return fmt.Errorf("user is a brokie")
	}

	// Update user balance
	updateBalanceQuery := `UPDATE user_data SET balance = balance - ? WHERE rowid = 1`
	_, err = tx.Exec(updateBalanceQuery, cost)
	if err != nil {
		return err
	}

	return nil
}

func SellCrypto(code string, price float64, sell_quantity float64) error {
	db, err := config.ConnectDB()
	if err != nil {
		return err
	}
	defer db.Close()

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

	var invested float64
	var crypto_count float64
	cryptoInfoQuery := `SELECT invested, crypto_count FROM crypto WHERE code = ?`
	err = tx.QueryRow(cryptoInfoQuery, code).Scan(&invested, &crypto_count)
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
	new_invested := invested - (price * sell_quantity)

	if new_crypto_count == 0 {
		// Delete the crypto entry if the new count is zero
		deleteCryptoQuery := `DELETE FROM crypto WHERE code = ?`
		_, err = tx.Exec(deleteCryptoQuery, code)
		if err != nil {
			return err
		}
	} else {
		// Update crypto holdings
		updateCryptoQuery := `UPDATE crypto SET invested = ?, crypto_count = ? WHERE code = ?`
		_, err = tx.Exec(updateCryptoQuery, new_invested, new_crypto_count, code)
		if err != nil {
			return err
		}
	}

	// Update user balance
	updateBalanceQuery := `UPDATE user_data SET balance = balance + ? WHERE rowid = 1`
	_, err = tx.Exec(updateBalanceQuery, price*sell_quantity)
	if err != nil {
		return err
	}

	return nil
}
