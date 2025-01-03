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

// Test with no crypto entries
func TestGetCryptoPortfolio1(t *testing.T) {

	type CryptoData struct {
		Code         string  `json:"code"`
		Invested     float64 `json:"invested"`
		Crypto_count float64 `json:"crypto_count"`
	}

	var expected []CryptoData

	crypto_items, err := GetCryptoPortfolio(testDB)
	if err != nil {
		t.Fatalf("Failed to retrieve crypto portfolio (1): %v", err)
	}
	assert.EqualValues(t, expected, crypto_items, "Crypto data should not exist currently")
}

// Test with no stock entries
func TestGetStockPortfolio1(t *testing.T) {

	type StockData struct {
		Code        string  `json:"code"`
		Invested    float64 `json:"invested"`
		Stock_count float64 `json:"stock_count"`
	}

	var expected []StockData

	stock_items, err := GetStockPortfolio(testDB)
	if err != nil {
		t.Fatalf("Failed to retrieve stock portfolio (1): %v", err)
	}
	assert.EqualValues(t, expected, stock_items, "Stock data should not exist currently")
}

func TestBuyCrypto(t *testing.T) {
	err := BuyCrypto(testDB, "TEST", 10, 5)

	// User has no money so the first one should fail
	assert.Equal(t, "insufficient funds", err.Error(), "User should have 0 balance")

	// Make sure no crypto was added
	var count int
	testDB.QueryRow(`SELECT COUNT(*) FROM crypto WHERE code = TEST`).Scan(&count)
	if count != 0 {
		t.Fatalf("Crypto purchased with insufficient funds")
	}

	// Give User money
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
		Invested    float64 `json:"invested"`
		CryptoCount float64 `json:"crypto_count"`
	}
	var crypto_data CryptoData
	err = testDB.QueryRow(`SELECT invested, crypto_count FROM crypto WHERE code = "TEST"`).Scan(&crypto_data.Invested, &crypto_data.CryptoCount)
	if err != nil {
		t.Fatalf("Error retrieving crypto: %v", err)
	}

	if crypto_data.CryptoCount != 5 {
		t.Fatalf("Incorrect crypto count. Expected 5, Got %f", crypto_data.CryptoCount)
	}
}
