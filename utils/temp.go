package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
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

func Fetch_api(symbol string) (float64, error) {
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		return 0, fmt.Errorf("API key is not set")
	}

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
	//fmt.Printf("Stock Quote for %s:\n", symbol)
	//fmt.Printf("Current Price: $%.2f\n", quote.CurrentPrice)
	//fmt.Printf("Change: $%.2f\n", quote.Change)
	//fmt.Printf("Percentage Change: %.2f%%\n", quote.PercentageChange)
	//fmt.Printf("High Price: $%.2f\n", quote.High)
	//fmt.Printf("Low Price: $%.2f\n", quote.Low)
	//fmt.Printf("Open Price: $%.2f\n", quote.Open)
	//fmt.Printf("Previous Close: $%.2f\n", quote.PreviousClose)
	//timestamp := time.Unix(quote.Timestamp, 0)
	//fmt.Println("Timestamp:", timestamp.Format(time.RFC3339))

	// Returning only the stock price for database calculating the profit gain and loss
	return quote.CurrentPrice, nil
}

func Fetch_timestamp(symbol string) (time.Time, error) {
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		return time.Now(), fmt.Errorf("API key is not set")
	}

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
	timestamp := time.Unix(quote.Timestamp, 0)
	return timestamp, nil
}

func Fetch_previous_close(symbol string) (float64, error) {
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		return 0, fmt.Errorf("API key is not set")
	}

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
	return quote.PreviousClose, nil
}

func Fetch_open_price(symbol string) (float64, error) {
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		return 0, fmt.Errorf("API key is not set")
	}

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
	return quote.Open, nil
}

func Fetch_low_price(symbol string) (float64, error) {
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		return 0, fmt.Errorf("API key is not set")
	}

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
	return quote.Low, nil
}

func Fetch_high_price(symbol string) (float64, error) {
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		return 0, fmt.Errorf("API key is not set")
	}

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
	return quote.High, nil
}

func Fetch_percent_change(symbol string) (float64, error) {
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		return 0, fmt.Errorf("API key is not set")
	}

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
	return quote.PercentageChange, nil
}

func Fetch_change(symbol string) (float64, error) {
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		return 0, fmt.Errorf("API key is not set")
	}

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
	return quote.Change, nil
}
