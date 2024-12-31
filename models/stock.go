package models

import (
	"database/sql"
	"github.com/iamloganwalsh/stock-trading-simulator/config"
	"fmt"
)

func BuyStock(code string, cost float64, stock_count float64) error {
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
	var old_count float64
	query := `SELECT invested, stock_count FROM stock WHERE code = ?`
	err = tx.QueryRow(query, code).Scan(&invested, &old_count)
	if err != nil {
		if err == sql.ErrNoRows { // Add new entry
			addQuery := `INSERT INTO stock (code, invested, stock_count) VALUES (?, ?, ?)`
			_, err = tx.Exec(addQuery, code, cost, stock_count)
			if err != nil {
				return err
			}
		} else {
			return err
		}
	} else {
		// Update existing entry
		updateQuery := `UPDATE stock SET invested = ?, stock_count = ? WHERE code = ?`
		_, err = tx.Exec(updateQuery, invested+cost, stock_count+old_count, code)
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


func SellStock(code string, price float64, sell_quantity float64) error {
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
    var stock_count float64
    stockInfoQuery := `SELECT invested, stock_count FROM stock WHERE code = ?`
    err = tx.QueryRow(stockInfoQuery, code).Scan(&invested, &stock_count)
    if err != nil {
        if err == sql.ErrNoRows {
            return fmt.Errorf("user does not own specified stock")
        }
        return err
    }

    if sell_quantity > stock_count {
        return fmt.Errorf("user does not own that much stock; current amount owned: %.6f", stock_count)
    }

    // Calculate new stock count and invested amount
    new_stock_count := stock_count - sell_quantity
    new_invested := invested * (new_stock_count / stock_count)

    if new_stock_count == 0 {
        // Delete the stock entry if the new count is zero
        deleteStockQuery := `DELETE FROM stock WHERE code = ?`
        _, err = tx.Exec(deleteStockQuery, code)
        if err != nil {
            return err
        }
    } else {
        // Update stock holdings
        updateStockQuery := `UPDATE stock SET invested = ?, stock_count = ? WHERE code = ?`
        _, err = tx.Exec(updateStockQuery, new_invested, new_stock_count, code)
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

