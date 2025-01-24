package utils

import (
	"context"
	"testing"
	"time"
)

func TestRedisClient(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 75*time.Second)
	defer cancel()

	t.Run("TestNewRedisClient", func(t *testing.T) {
		client, err := NewRedisClient(ctx, "localhost:6379", "", 0)
		if err != nil {
			t.Fatalf("Failed to connect to Redis: %v\nMake sure Redis is running on localhost:6379", err)
		}
		defer client.Close()

		if client == nil {
			t.Fatal("Expected non-nil client")
		}
	})

	// Creating a client for subsequent tests
	client, err := NewRedisClient(ctx, "localhost:6379", "", 0)
	if err != nil {
		t.Fatalf("Failed to setup Redis client for tests: %v", err)
	}
	defer client.Close()

	t.Run("TestStockPriceCache", func(t *testing.T) {
		symbol := "AAPL"
		price := 150.50

		// Test setting cache
		err := client.CacheStockPrice(symbol, price)
		if err != nil {
			t.Fatalf("Failed to cache stock price: %v", err)
		}

		// Test getting cache
		cachedPrice, err := client.GetCacheStockQuote(symbol)
		if err != nil {
			t.Fatalf("Failed to get cached stock price: %v", err)
		}
		if cachedPrice != price {
			t.Errorf("Expected price %v, got %v", price, cachedPrice)
		}
	})

	t.Run("TestCryptoPriceCache", func(t *testing.T) {
		symbol := "BTC"
		price := 45000.75

		// Test setting cache
		err := client.CacheCryptoPrice(symbol, price)
		if err != nil {
			t.Fatalf("Failed to cache crypto price: %v", err)
		}

		// Test getting cache
		cachedPrice, err := client.GetCacheCryptoQuote(symbol)
		if err != nil {
			t.Fatalf("Failed to get cached crypto price: %v", err)
		}
		if cachedPrice != price {
			t.Errorf("Expected price %v, got %v", price, cachedPrice)
		}
	})

	t.Run("TestNonExistentCache", func(t *testing.T) {
		symbol := "NONEXISTENT"

		price, err := client.GetCacheStockQuote(symbol)
		if err != nil {
			t.Errorf("Expected no error for non-existent stock, got: %v", err)
		}
		if price != 0 {
			t.Errorf("Expected 0 for non-existent stock, got %v", price)
		}
	})

	t.Run("TestCacheExpiration", func(t *testing.T) {
		symbol := "MSFT"
		price := 300.75

		err := client.CacheStockPrice(symbol, price)
		if err != nil {
			t.Fatalf("Failed to cache stock price: %v", err)
		}

		// Verify immediate retrieval
		cachedPrice, err := client.GetCacheStockQuote(symbol)
		if err != nil {
			t.Fatalf("Failed to get cached stock price: %v", err)
		}
		if cachedPrice != price {
			t.Errorf("Expected price %v, got %v", price, cachedPrice)
		}

		time.Sleep(65 * time.Second)

		expiredPrice, err := client.GetCacheStockQuote(symbol)

		if err != nil {
			t.Fatalf("Failed to get expired stock price: %v", err)
		}
		if expiredPrice != 0 {
			t.Errorf("Expected ecpired to return 0, got %v", expiredPrice)
		}
	})
}
