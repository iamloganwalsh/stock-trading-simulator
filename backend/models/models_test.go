package models

import (
	"database/sql"
	"os"
	"testing"

	"github.com/iamloganwalsh/stock-trading-simulator/config"
	"github.com/stretchr/testify/assert"
)

var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error
	testDB, err = config.ConnectDB("./tests.db")
	if err != nil {
		panic("Failed to connect to test database: " + err.Error())
	}

	err = config.InitDB(testDB)
	if err != nil {
		panic("Failed to initialize testing database: " + err.Error())
	}

	code := m.Run()

	testDB.Close()
	//os.Remove("./tests.db")
	os.Exit(code)
}

// user.go
func TestInitUser(t *testing.T) {

	username := "test_username"
	err := InitUser(testDB, username)
	assert.NoError(t, err, "Failed to init user")

	// Check details are correct
	type User struct {
		Username    string  `json:"username"`
		Balance     float64 `json:"balance"`
		Profit_loss float64 `json:"profit_loss"`
	}

	var user User

	err = testDB.QueryRow("SELECT username, balance, profit_loss FROM user_data").Scan(&user.Username, &user.Balance, &user.Profit_loss)

	if err != nil {
		t.Fatalf("Failed to retrieve user data: %v", err)
	}

	assert.Equal(t, "test_username", user.Username, "Username does not match")
	assert.Equal(t, float64(0), user.Balance, "Balance is not 0")
	assert.Equal(t, float64(0), user.Profit_loss, "Profit Loss is not 0")
}

func TestGetUsername(t *testing.T) {
	username, err := GetUsername(testDB)
	if err != nil {
		t.Fatalf("Failed to retrieve username: %v", err)
	}
	assert.Equal(t, "test_username", username, "Username is incorrect")
}

func TestGetBalance(t *testing.T) {
	balance, err := GetBalance(testDB)
	if err != nil {
		t.Fatalf("Failed to retrieve balance: %v", err)
	}
	assert.Equal(t, float64(0), balance, "Balance should be 0")
}

func TestGetProfitLoss(t *testing.T) {
	profit_loss, err := GetProfitLoss(testDB)
	if err != nil {
		t.Fatalf("Failed to retrieve profit loss: %v", err)
	}
	assert.Equal(t, float64(0), profit_loss, "Profit Loss should be 0")
}

func TestGetCryptoPortfolio(t *testing.T) {

	type CryptoData struct {
		Code         string  `json:"code"`
		Crypto_count float64 `json:"crypto_count"`
	}

	var expected []CryptoData

	crypto_items, err := GetCryptoPortfolio(testDB)
	if err != nil {
		t.Fatalf("Failed to retrieve crypto portfolio: %v", err)
	}
	assert.EqualValues(t, expected, crypto_items, "Crypto data should not exist currently")

	testDB.Exec("INSERT INTO crypto (code, crypto_count) VALUES (?, ?)", "TEST1", 20)
	testDB.Exec("INSERT INTO crypto (code, crypto_count) VALUES (?, ?)", "TEST2", 10)

	crypto_items, err = GetCryptoPortfolio(testDB)
	if err != nil {
		t.Fatalf("Failed to retrieve crypto portfolio: %v", err)
	}

	expected = []CryptoData{
		{Code: "TEST1", Crypto_count: 20},
		{Code: "TEST2", Crypto_count: 10},
	}

	assert.EqualValues(t, expected, crypto_items, "Crypto data does not match expected values")

}

func TestGetStockPortfolio(t *testing.T) {

	type StockData struct {
		Code        string  `json:"code"`
		Stock_count float64 `json:"stock_count"`
	}

	var expected []StockData

	stock_items, err := GetStockPortfolio(testDB)
	if err != nil {
		t.Fatalf("Failed to retrieve stock portfolio: %v", err)
	}
	assert.EqualValues(t, expected, stock_items, "Stock data should not exist currently")

	testDB.Exec("INSERT INTO stock (code, stock_count) VALUES (?, ?)", "TEST1", 20)
	testDB.Exec("INSERT INTO stock (code, stock_count) VALUES (?, ?)", "TEST2", 10)

	stock_items, err = GetStockPortfolio(testDB)
	if err != nil {
		t.Fatalf("Failed to retrieve stock portfolio: %v", err)
	}

	expected = []StockData{
		{Code: "TEST1", Stock_count: 20},
		{Code: "TEST2", Stock_count: 10},
	}

	assert.EqualValues(t, expected, stock_items, "Stock data does not match expected values")

}

// crypto.go
func TestBuyCrypto(t *testing.T) {

	// Set balance to 0
	testDB.Exec(`UPDATE user_data SET balance = 0 WHERE rowid = 1`)

	err := BuyCrypto(testDB, "TEST", 10, 5)

	// User has no money so the first one should fail
	if err == nil {
		t.Fatalf("User purchased crypto without funds")
	}

	// Make sure no crypto was added
	var count int
	testDB.QueryRow(`SELECT COUNT(*) FROM crypto WHERE code = TEST`).Scan(&count)
	if count != 0 {
		t.Fatalf("Crypto purchased with insufficient funds")
	}

	// Give user money
	testDB.Exec(`UPDATE user_data SET balance = 100 WHERE rowid = 1`)

	// Attempt purchase again
	err = BuyCrypto(testDB, "TEST", 10, 5)

	// Test for successful purchase
	if err != nil {
		t.Fatalf("Purchase failed: %v", err)
	}

	// User balance correctly deducted
	var user_balance float64
	err = testDB.QueryRow(`SELECT balance FROM user_data`).Scan(&user_balance)
	if err != nil {
		t.Fatalf("Error retrieving user balance: %v", err)
	}
	if user_balance != float64(50) {
		t.Fatalf("Incorrect user balance after purchase. Expected %f, Got %f", float64(50), user_balance)
	}

	// Crypto successfuly purchased
	type CryptoData struct {
		Code        string  `json:"code"`
		CryptoCount float64 `json:"crypto_count"`
	}
	var crypto_data CryptoData
	err = testDB.QueryRow(`SELECT crypto_count FROM crypto WHERE code = "TEST"`).Scan(&crypto_data.CryptoCount)
	if err != nil {
		t.Fatalf("Error retrieving crypto: %v", err)
	}

	if crypto_data.CryptoCount != 5 {
		t.Fatalf("Incorrect crypto count. Expected 5, Got %f", crypto_data.CryptoCount)
	}

	// Purchasing more should add to existing entry
	_, err = testDB.Exec(`UPDATE user_data SET balance = balance + 20 WHERE rowid = 1`)
	if err != nil {
		t.Fatalf("Error updating user balance: %v", err)
	}
	err = BuyCrypto(testDB, "TEST", 10, 2)
	if err != nil {
		t.Fatalf("Error purchasing crypto: %v", err)
	}

	// Test database entry
	err = testDB.QueryRow(`SELECT crypto_count FROM crypto WHERE code = "TEST"`).Scan(&crypto_data.CryptoCount)
	if err != nil {
		t.Fatalf("Error retrieving crypto count: %v", err)
	}
	assert.Equal(t, float64(7), crypto_data.CryptoCount, "Second purchase not update crypto count correctly")
}

func TestSellCrypto(t *testing.T) {
	// Sell non existant crypto
	err := SellCrypto(testDB, "FAKE", 10, 5)
	if err == nil {
		t.Fatalf("Crypto should not exist")
	}

	// Test user balance is still 50
	var user_balance float64
	err = testDB.QueryRow(`SELECT balance FROM user_data`).Scan(&user_balance)
	if err != nil {
		t.Fatalf("Error retrieving user balance: %v", err)
	}
	assert.Equal(t, float64(50), user_balance, "Incorrect balance")

	err = SellCrypto(testDB, "TEST", 10, 4) // Sell 4 crypto for 10 each
	if err != nil {
		t.Fatalf("Error selling crypto: %v", err)
	}

	// Test user balance is now 70
	err = testDB.QueryRow(`SELECT balance FROM user_data`).Scan(&user_balance)
	if err != nil {
		t.Fatalf("Error retrieving user balance: %v", err)
	}
	assert.Equal(t, float64(90), user_balance, "Incorrect balance")

	// Test 2 tokens were deducted
	var crypto_count float64
	err = testDB.QueryRow(`SELECT crypto_count FROM crypto WHERE code == "TEST"`).Scan(&crypto_count)
	if err != nil {
		t.Fatalf("Error retrieving crypto count: %v", err)
	}
	assert.Equal(t, float64(3), crypto_count, "Incorrect crypto count")

	// Test selling too many tokens
	err = SellCrypto(testDB, "TEST", 10, 4)
	if err == nil {
		t.Fatalf("Sold more crypto than owned. Owned 3, successfuly sold 4.")
	}

	// Make sure balance and crypto count doesnt change
	err = testDB.QueryRow(`SELECT balance FROM user_data`).Scan(&user_balance)
	if err != nil {
		t.Fatalf("Error retrieving user balance: %v", err)
	}
	assert.Equal(t, float64(90), user_balance, "Incorrect balance")
	err = testDB.QueryRow(`SELECT crypto_count FROM crypto WHERE code == "TEST"`).Scan(&crypto_count)
	if err != nil {
		t.Fatalf("Error retrieving crypto count: %v", err)
	}
	assert.Equal(t, float64(3), crypto_count, "Incorrect crypto count")

	// Test selling all and deleting from database
	err = SellCrypto(testDB, "TEST", 10, 3)
	if err != nil {
		t.Fatalf("Error selling crypto: %v", err)
	}

	// Make sure database correctly updated
	err = testDB.QueryRow(`SELECT balance FROM user_data`).Scan(&user_balance)
	if err != nil {
		t.Fatalf("Error retrieving user balance: %v", err)
	}
	assert.Equal(t, float64(120), user_balance, "Incorrect balance")
	err = testDB.QueryRow(`SELECT crypto_count FROM crypto WHERE code == "TEST"`).Scan(&crypto_count)
	if err != sql.ErrNoRows {
		t.Fatalf("Database entry not deleted: %v", err)
	}
}

// stock.go
func TestBuyStock(t *testing.T) {

	// Set balance to 0
	testDB.Exec(`UPDATE user_data SET balance = 0 WHERE rowid = 1`)

	err := BuyStock(testDB, "TEST", 10, 5)

	// User has no money so the first one should fail
	if err == nil {
		t.Fatalf("User purchased stock without funds")
	}

	// Make sure no stock was added
	var count int
	testDB.QueryRow(`SELECT COUNT(*) FROM stock WHERE code = TEST`).Scan(&count)
	if count != 0 {
		t.Fatalf("Stock purchased with insufficient funds")
	}

	// Give user money
	testDB.Exec(`UPDATE user_data SET balance = 100 WHERE rowid = 1`)

	// Attempt purchase again
	err = BuyStock(testDB, "TEST", 10, 5)

	// Test for successful purchase
	if err != nil {
		t.Fatalf("Purchase failed: %v", err)
	}

	// User balance correctly deducted
	var user_balance float64
	err = testDB.QueryRow(`SELECT balance FROM user_data`).Scan(&user_balance)
	if err != nil {
		t.Fatalf("Error retrieving user balance: %v", err)
	}
	if user_balance != float64(50) {
		t.Fatalf("Incorrect user balance after purchase. Expected %f, Got %f", float64(50), user_balance)
	}

	// Stock successfuly purchased
	type StockData struct {
		Code       string  `json:"code"`
		StockCount float64 `json:"stock_count"`
	}
	var stock_data StockData
	err = testDB.QueryRow(`SELECT stock_count FROM stock WHERE code = "TEST"`).Scan(&stock_data.StockCount)
	if err != nil {
		t.Fatalf("Error retrieving stock: %v", err)
	}

	if stock_data.StockCount != 5 {
		t.Fatalf("Incorrect stock count. Expected 5, Got %f", stock_data.StockCount)
	}

	// Purchasing more should add to existing entry
	_, err = testDB.Exec(`UPDATE user_data SET balance = balance + 20 WHERE rowid = 1`)
	if err != nil {
		t.Fatalf("Error updating user balance: %v", err)
	}
	err = BuyStock(testDB, "TEST", 10, 2)
	if err != nil {
		t.Fatalf("Error purchasing stock: %v", err)
	}

	// Test database entry
	err = testDB.QueryRow(`SELECT stock_count FROM stock WHERE code = "TEST"`).Scan(&stock_data.StockCount)
	if err != nil {
		t.Fatalf("Error retrieving stock count: %v", err)
	}
	assert.Equal(t, float64(7), stock_data.StockCount, "Second purchase not updating stock count correctly")
}

func TestSellStock(t *testing.T) {
	// Sell non existant stock
	err := SellStock(testDB, "FAKE", 10, 5)
	if err == nil {
		t.Fatalf("Stock should not exist")
	}

	// Test user balance is still 50
	var user_balance float64
	err = testDB.QueryRow(`SELECT balance FROM user_data`).Scan(&user_balance)
	if err != nil {
		t.Fatalf("Error retrieving user balance: %v", err)
	}
	assert.Equal(t, float64(50), user_balance, "Incorrect balance")

	err = SellStock(testDB, "TEST", 10, 4) // Sell 4 stock for 10 each
	if err != nil {
		t.Fatalf("Error selling stock: %v", err)
	}

	// Test user balance is now 70
	err = testDB.QueryRow(`SELECT balance FROM user_data`).Scan(&user_balance)
	if err != nil {
		t.Fatalf("Error retrieving user balance: %v", err)
	}
	assert.Equal(t, float64(90), user_balance, "Incorrect balance")

	// Test 2 tokens were deducted
	var stock_count float64
	err = testDB.QueryRow(`SELECT stock_count FROM stock WHERE code == "TEST"`).Scan(&stock_count)
	if err != nil {
		t.Fatalf("Error retrieving stock count: %v", err)
	}
	assert.Equal(t, float64(3), stock_count, "Incorrect stock count")

	// Test selling too many tokens
	err = SellStock(testDB, "TEST", 10, 4)
	if err == nil {
		t.Fatalf("Sold more stock than owned. Owned 3, successfuly sold 4.")
	}

	// Make sure balance and stock count doesnt change
	err = testDB.QueryRow(`SELECT balance FROM user_data`).Scan(&user_balance)
	if err != nil {
		t.Fatalf("Error retrieving user balance: %v", err)
	}
	assert.Equal(t, float64(90), user_balance, "Incorrect balance")
	err = testDB.QueryRow(`SELECT stock_count FROM stock WHERE code == "TEST"`).Scan(&stock_count)
	if err != nil {
		t.Fatalf("Error retrieving stock count: %v", err)
	}
	assert.Equal(t, float64(3), stock_count, "Incorrect stock count")

	// Test selling all and deleting from database
	err = SellStock(testDB, "TEST", 10, 3)
	if err != nil {
		t.Fatalf("Error selling stock: %v", err)
	}

	// Make sure database correctly updated
	err = testDB.QueryRow(`SELECT balance FROM user_data`).Scan(&user_balance)
	if err != nil {
		t.Fatalf("Error retrieving user balance: %v", err)
	}
	assert.Equal(t, float64(120), user_balance, "Incorrect balance")
	err = testDB.QueryRow(`SELECT stock_count FROM stock WHERE code == "TEST"`).Scan(&stock_count)
	if err != sql.ErrNoRows {
		t.Fatalf("Database entry not deleted: %v", err)
	}
}
