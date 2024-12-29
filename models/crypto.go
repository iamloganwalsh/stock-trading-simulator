package models

import (
	"database/sql"
	"github.com/iamloganwalsh/stock-trading-simulator/config"
	"fmt"
)

func BuyCrypto(code string, cost float64, cryptoCount float64) error {
	// Start a new transaction
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
	var oldCount float64
	query := `SELECT invested, crypto_count FROM crypto WHERE code = ?`
	err = tx.QueryRow(query, code).Scan(&invested, &oldCount)
	if err != nil {
		if err == sql.ErrNoRows { // Add new entry
			addQuery := `INSERT INTO crypto (code, invested, crypto_count) VALUES (?, ?, ?)`
			_, err = tx.Exec(addQuery, code, cost, cryptoCount)
			if err != nil {
				return err
			}
		} else {
			return err
		}
	} else {
		// Update existing entry
		updateQuery := `UPDATE crypto SET invested = ?, crypto_count = ? WHERE code = ?`
		_, err = tx.Exec(updateQuery, invested+cost, cryptoCount+oldCount, code)
		if err != nil {
			return err
		}
	}

	// Double check that user can afford	
	userBalance, err := GetBalance() // From user.go
	if err != nil {
		return err
	}
	if userBalance < cost {
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


func SellCrypto(code string, price float64, sellQuantity float64) error {
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
    var cryptoCount float64
    cryptoInfoQuery := `SELECT invested, crypto_count FROM crypto WHERE code = ?`
    err = tx.QueryRow(cryptoInfoQuery, code).Scan(&invested, &cryptoCount)
    if err != nil {
        if err == sql.ErrNoRows {
            return fmt.Errorf("user does not own specified crypto")
        }
        return err
    }

    if sellQuantity > cryptoCount {
        return fmt.Errorf("user does not own that much crypto; current amount owned: %.6f", cryptoCount)
    }

    // Calculate new crypto count and invested amount
    newCryptoCount := cryptoCount - sellQuantity
    newInvested := invested * (newCryptoCount / cryptoCount)

    if newCryptoCount == 0 {
        // Delete the crypto entry if the new count is zero
        deleteCryptoQuery := `DELETE FROM crypto WHERE code = ?`
        _, err = tx.Exec(deleteCryptoQuery, code)
        if err != nil {
            return err
        }
    } else {
        // Update crypto holdings
        updateCryptoQuery := `UPDATE crypto SET invested = ?, crypto_count = ? WHERE code = ?`
        _, err = tx.Exec(updateCryptoQuery, newInvested, newCryptoCount, code)
        if err != nil {
            return err
        }
    }

    // Update user balance
    updateBalanceQuery := `UPDATE user_data SET balance = balance + ? WHERE rowid = 1`
    _, err = tx.Exec(updateBalanceQuery, price)
    if err != nil {
        return err
    }

    return nil
}
