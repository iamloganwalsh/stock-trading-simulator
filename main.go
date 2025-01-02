package main

import (
	"log"
	"net/http"

	"github.com/iamloganwalsh/stock-trading-simulator/config"
	//"github.com/iamloganwalsh/stock-trading-simulator/utils"
	"github.com/joho/godotenv"
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

	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//utils.Fetch_api()

	log.Println("Starting server on localhost:3000...")
	log.Fatal(http.ListenAndServe("localhost:3000", nil))
}
