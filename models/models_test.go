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
