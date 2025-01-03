package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/iamloganwalsh/stock-trading-simulator/config"
	"github.com/iamloganwalsh/stock-trading-simulator/routes"
	//"github.com/iamloganwalsh/stock-trading-simulator/utils"
)

// Struct to parse the response for stock quotes
type StockQuote struct {
	CurrentPrice     float64 `json:"c"`
	Change           float64 `json:"d"`
	PercentageChange float64 `json:"dp"`
	High             float64 `json:"h"`
	Low              float64 `json:"l"`
	Open             float64 `json:"o"`
	PreviousClose    float64 `json:"pc"`
	Timestamp        int64   `json:"t"` // Unix timestamp
}

func main() {

	// Database Connect or Init
	db, err := config.ConnectDB()

	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	err = config.InitDB(db)
	if err != nil {
		log.Fatalf("Failed to initialize the database: %v", err)
	}

	router := mux.NewRouter()

	// User routes
	router.HandleFunc("/user/create", routes.InitUserHandler).Methods("POST")
	router.HandleFunc("/user/username", routes.GetUsernameHandler).Methods("GET")
	router.HandleFunc("/user/balance", routes.GetBalanceHandler).Methods("GET")
	router.HandleFunc("/user/profit_loss", routes.GetProfitLossHandler).Methods("GET")
	router.HandleFunc("/user/crypto_portfolio", routes.GetCryptoPortfolioHandler).Methods("GET")
	router.HandleFunc("/user/stock_portfolio", routes.GetStockPortfolioHandler).Methods("GET")

	// Trade routes (Crypto & Stock)
	router.HandleFunc("/crypto/buy", routes.BuyCryptoHandler).Methods("POST")
	router.HandleFunc("/crypto/sell", routes.SellCryptoHandler).Methods("POST") // Could DELETE or UPDATE so POST for versatility
	router.HandleFunc("/stock/buy", routes.BuyStockHandler).Methods("POST")
	router.HandleFunc("/stock/sell", routes.SellStockHandler).Methods("POST") // Could DELETE or UPDATE so POST for versatility

	log.Println("Starting server on localhost:3000...")
	log.Fatal(http.ListenAndServe("localhost:3000", router))
}
