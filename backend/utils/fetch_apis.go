package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

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

var (
	redisClient *RedisClient
	ctx         context.Context
)

func init() {
	var err error
	ctx = context.Background()
	redisClient, err = NewRedisClient(ctx, "localhost:6379", "", 0)
	if err != nil {
		log.Printf("Failed to connect to Redis: %v", err)
	}
}

func Fetch_api(symbol string) (*StockQuote, error) {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("API key is not set")
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

	return &quote, nil
}

// needs more development!!!

func Fetch_stock_price(symbol string) (float64, error) {
	if redisClient != nil {
		price, err := redisClient.GetCacheStockQuote(symbol)
		if err == nil && price != 0 {
			return price, nil
		}
	}
	quote, err := Fetch_api(symbol)
	if err != nil {
		return 0, err
	}

	if redisClient != nil {
		if err := redisClient.CacheStockPrice(symbol, quote.CurrentPrice); err != nil {
			log.Printf("Failed to cache stock price: %v", err)
		}
	}

	return quote.CurrentPrice, nil
}

func Fetch_crypto_price(symbol string) (float64, error) {
	if redisClient != nil {
		price, err := redisClient.GetCacheCryptoQuote(symbol)
		if err == nil && price != 0 {
			return price, nil
		}
	}
	quote, err := Fetch_api(symbol)
	if err != nil {
		return 0, err
	}

	if redisClient != nil {
		if err := redisClient.CacheCryptoPrice(symbol, quote.CurrentPrice); err != nil {
			log.Printf("Failed to cache crypto price: %v", err)
		}
	}

	return quote.CurrentPrice, nil
}

func Fetch_timestamp(symbol string) (time.Time, error) {
	quote, err := Fetch_api(symbol)

	if err != nil {
		return time.Time{}, err
	}

	return time.Unix(quote.Timestamp, 0), nil
}

func Fetch_previous_close(symbol string) (float64, error) {
	quote, err := Fetch_api(symbol)

	if err != nil {
		return 0, err
	}

	return quote.PreviousClose, nil
}

func Fetch_open_price(symbol string) (float64, error) {
	quote, err := Fetch_api(symbol)

	if err != nil {
		return 0, err
	}

	return quote.Open, nil
}

func Fetch_low_price(symbol string) (float64, error) {
	quote, err := Fetch_api(symbol)

	if err != nil {
		return 0, err
	}

	return quote.Low, nil
}

func Fetch_high_price(symbol string) (float64, error) {
	quote, err := Fetch_api(symbol)

	if err != nil {
		return 0, err
	}

	return quote.High, nil
}

func Fetch_percent_change(symbol string) (float64, error) {
	quote, err := Fetch_api(symbol)

	if err != nil {
		return 0, err
	}

	return quote.PercentageChange, nil
}

func Fetch_change(symbol string) (float64, error) {
	quote, err := Fetch_api(symbol)

	if err != nil {
		return 0, err
	}

	return quote.Change, nil
}

type YahooFinanceResponse struct {
	Chart struct {
		Result []struct {
			Timestamp  []int64 `json:"timestamp"`
			Indicators struct {
				Quote []struct {
					Close []float64 `json:"close"`
				} `json:"quote"`
			} `json:"indicators"`
		} `json:"result"`
	} `json:"chart"`
}

func Fetch_prev_crypto(symbol string) ([]float64, error) {
	// Define the Yahoo Finance URL
	url := fmt.Sprintf("https://query1.finance.yahoo.com/v7/finance/chart/%s?interval=1m&range=1d", symbol)

	// Make the HTTP request
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch data: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	// Parse the JSON response
	var data YahooFinanceResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %v", err)
	}

	// Extract the prices for the last 20 minutes
	if len(data.Chart.Result) == 0 || len(data.Chart.Result[0].Indicators.Quote) == 0 {
		return nil, fmt.Errorf("no data found")
	}

	prices := data.Chart.Result[0].Indicators.Quote[0].Close
	timestamps := data.Chart.Result[0].Timestamp

	// Filter the last 20 minutes of data
	var recentPrices []float64
	currentTime := time.Now().Unix()
	for i, timestamp := range timestamps {
		if currentTime-timestamp <= 20*60 {
			recentPrices = append(recentPrices, prices[i])
		}
	}

	return recentPrices, nil
}
