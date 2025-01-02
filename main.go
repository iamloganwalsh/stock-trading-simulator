package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/iamloganwalsh/stock-trading-simulator/config"
	"github.com/iamloganwalsh/stock-trading-simulator/routes"
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

	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		fmt.Println("API key is not set")
		return
	}

	symbol := "AAPL" // Example symbol for Apple (AAPL), Microsoft (MSFT), Meta (META)

	// Construct the API URL
	url := fmt.Sprintf("https://finnhub.io/api/v1/quote?symbol=%s&token=%s", symbol, apiKey)

	// Send GET request to the API
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Failed to fetch data: %v", err)
	}
	defer resp.Body.Close()

	// Check if the response status code is OK
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Error: Received status code %d", resp.StatusCode)
	}

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}

	// Parse the JSON response into the StockQuote struct
	var quote StockQuote
	err = json.Unmarshal(body, &quote)
	if err != nil {
		log.Fatalf("Failed to parse JSON: %v", err)
	}

	// Print the stock quote details
	fmt.Printf("Stock Quote for %s:\n", symbol)
	fmt.Printf("Current Price: $%.2f\n", quote.CurrentPrice)
	fmt.Printf("Change: $%.2f\n", quote.Change)
	fmt.Printf("Percentage Change: %.2f%%\n", quote.PercentageChange)
	fmt.Printf("High Price: $%.2f\n", quote.High)
	fmt.Printf("Low Price: $%.2f\n", quote.Low)
	fmt.Printf("Open Price: $%.2f\n", quote.Open)
	fmt.Printf("Previous Close: $%.2f\n", quote.PreviousClose)
	timestamp := time.Unix(quote.Timestamp, 0)
	fmt.Println("Timestamp:", timestamp.Format(time.RFC3339))

	router := mux.NewRouter()
	router.HandleFunc("/user/create", routes.InitUserHandler).Methods("POST")
	router.HandleFunc("/user/username", routes.GetUsernameHandler).Methods("GET")

	// test these
	router.HandleFunc("/user/balance", routes.GetBalanceHandler).Methods("GET")
	router.HandleFunc("/user/profit_loss", routes.GetProfitLossHandler).Methods("GET")

	log.Println("Starting server on localhost:3000...")
	log.Fatal(http.ListenAndServe("localhost:3000", router))
}
