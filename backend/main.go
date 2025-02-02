package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/iamloganwalsh/stock-trading-simulator/config"
	"github.com/iamloganwalsh/stock-trading-simulator/routes"
	"github.com/iamloganwalsh/stock-trading-simulator/utils"
	"github.com/rs/cors"
)

func main() {

	// Database Connect and Init
	db, err := config.ConnectDB()

	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	err = config.InitDB(db)
	if err != nil {
		log.Fatalf("Failed to initialize the database: %v", err)
	}

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	router := mux.NewRouter()

	// User routes
	router.HandleFunc("/user/create", routes.InitUserHandler).Methods("POST")
	router.HandleFunc("/user/username", routes.GetUsernameHandler).Methods("GET")
	router.HandleFunc("/user/balance", routes.GetBalanceHandler).Methods("GET")
	router.HandleFunc("/user/profit_loss", routes.GetProfitLossHandler).Methods("GET")
	router.HandleFunc("/user/crypto_portfolio", routes.GetCryptoPortfolioHandler).Methods("GET")
	router.HandleFunc("/user/stock_portfolio", routes.GetStockPortfolioHandler).Methods("GET")
	router.HandleFunc("/user/trade_history", routes.GetTradeHistoryHandler).Methods("GET")

	// Trade routes (Crypto & Stock)
	router.HandleFunc("/crypto/buy", routes.BuyCryptoHandler).Methods("POST")
	router.HandleFunc("/crypto/sell", routes.SellCryptoHandler).Methods("POST") // Could DELETE or UPDATE so POST for versatility
	router.HandleFunc("/stock/buy", routes.BuyStockHandler).Methods("POST")
	router.HandleFunc("/stock/sell", routes.SellStockHandler).Methods("POST") // Could DELETE or UPDATE so POST for versatility

	// Fetching routes
	router.HandleFunc("/crypto/fetch/{code}", routes.FetchCryptoHandler).Methods("GET")
	router.HandleFunc("/stock/fetch/{code}", routes.FetchStockHandler).Methods("GET")
	router.HandleFunc("/crypto/fetch_prev/{code}", routes.FetchCryptoPrevHandler).Methods("GET")
	router.HandleFunc("/stock/fetch_prev/{code}", routes.FetchStockPrevHandler).Methods("GET")

	stockPrice, err := utils.Fetch_stock_price("AAPL")
	if err != nil {
		fmt.Println("Error fetching stock data:", err)
	} else {
		fmt.Printf("Current stock price of AAPL: $%2.f\n", stockPrice)
	}

	cryptoPrice, err := utils.Fetch_crypto_price("BINANCE:BTCUSDT")
	if err != nil {
		fmt.Println("Error fetching crypto data:", err)
	} else {
		fmt.Printf("Current crypto price of BTC/USDT: $%2.f\n", cryptoPrice)
	}

	handler := c.Handler(router)
	log.Println("Starting server on localhost:3000...")
	log.Fatal(http.ListenAndServe("localhost:3000", handler))
}
