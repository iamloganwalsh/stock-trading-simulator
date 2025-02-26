package models

import (
	"database/sql"
	"fmt"
	"time"
)

func BuyStock(db *sql.DB, code string, cost float64, stock_count float64) error {
	// Start a new transaction
	cost *= stock_count

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

	var currentStockCount float64
	query := `SELECT stock_count FROM stock WHERE code = ?`
	err = tx.QueryRow(query, code).Scan(&currentStockCount)
	if err != nil {
		if err == sql.ErrNoRows { // Add new entry
			addQuery := `INSERT INTO stock (code, stock_count) VALUES (?, ?)`
			_, err = tx.Exec(addQuery, code, stock_count)
			if err != nil {
				return err
			}
		} else {
			return err
		}
	} else {
		// Update existing entry
		updateQuery := `UPDATE stock SET stock_count = ? WHERE code = ?`
		_, err = tx.Exec(updateQuery, stock_count+currentStockCount, code)
		if err != nil {
			return err
		}
	}

	// Update trade history
	tradeQuery := `INSERT INTO trade_history (type, code, method, cost, date) VALUES (?, ?, ?, ?, ?)`
	_, err = tx.Exec(tradeQuery, "stock", code, "buy", cost, time.Now().Format(time.RFC3339))
	if err != nil {
		return err
	}

	// Update user balance
	updateBalanceQuery := `UPDATE user_data SET balance = balance - ? WHERE rowid = 1`
	_, err = tx.Exec(updateBalanceQuery, cost)
	if err != nil {
		return err
	}

	// Update user investment
	updateInvestmentQuery := `UPDATE user_data SET investment = investment + ? WHERE rowid = 1`
	_, err = tx.Exec(updateInvestmentQuery, cost)
	if err != nil {
		return err
	}

	return nil
}

func SellStock(db *sql.DB, code string, price float64, sell_quantity float64) error {

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

	var stock_count float64
	stockInfoQuery := `SELECT stock_count FROM stock WHERE code = ?`
	err = tx.QueryRow(stockInfoQuery, code).Scan(&stock_count)
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

	if new_stock_count == 0 {
		// Delete the stock entry if the new count is zero
		deleteStockQuery := `DELETE FROM stock WHERE code = ?`
		_, err = tx.Exec(deleteStockQuery, code)
		if err != nil {
			return err
		}
	} else {
		// Update stock holdings
		updateStockQuery := `UPDATE stock SET stock_count = ? WHERE code = ?`
		_, err = tx.Exec(updateStockQuery, new_stock_count, code)
		if err != nil {
			return err
		}
	}

	// Update trade history
	totalRevenue := price * sell_quantity
	updateTradeQuery := `INSERT INTO trade_history (type, code, method, cost, date) VALUES (?, ?, ?, ?, ?)`
	_, err = tx.Exec(updateTradeQuery, "stock", code, "sell", totalRevenue, time.Now().Format(time.RFC3339))

	if err != nil {
		return err
	}

	// Update user balance
	updateBalanceQuery := `UPDATE user_data SET balance = balance + ? WHERE rowid = 1`
	_, err = tx.Exec(updateBalanceQuery, price*sell_quantity)
	if err != nil {
		return err
	}

	// Update user investment
	updateInvestmentQuery := `UPDATE user_data SET investment = investment - ? WHERE rowid = 1`
	_, err = tx.Exec(updateInvestmentQuery, price*sell_quantity)
	if err != nil {
		return err
	}

	return nil
}
