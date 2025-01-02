package routes

import (
	"encoding/json"
	"net/http"

	"github.com/iamloganwalsh/stock-trading-simulator/models"
)

func BuyCryptoHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method should be POST", http.StatusMethodNotAllowed)
		return
	}

	type CryptoDetails struct {
		Code   string  `json:"code"`
		Cost   float64 `json:"cost"`
		Amount float64 `json:"crypto_count"`
	}

	var userReq CryptoDetails
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&userReq)
	if err != nil {
		http.Error(w, "Bad Request: "+err.Error(), http.StatusBadRequest)
		return
	}

	err = models.BuyCrypto(userReq.Code, userReq.Cost, userReq.Amount)
	if err != nil {
		http.Error(w, "Internal server error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Crypto purchased successfully"))
}

func SellCryptoHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method should be POST", http.StatusMethodNotAllowed)
		return
	}

	type CryptoDetails struct {
		Code   string  `json:"code"`
		Cost   float64 `json:"cost"`
		Amount float64 `json:"crypto_count"`
	}

	var userReq CryptoDetails
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&userReq)
	if err != nil {
		http.Error(w, "Bad Request: "+err.Error(), http.StatusBadRequest)
		return
	}

	err = models.SellCrypto(userReq.Code, userReq.Cost, userReq.Amount)
	if err != nil {
		http.Error(w, "Internal server error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Crypto sold successfully"))
}

func BuyStockHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method should be POST", http.StatusMethodNotAllowed)
		return
	}

	type StockDetails struct {
		Code   string  `json:"code"`
		Cost   float64 `json:"cost"`
		Amount float64 `json:"stock_count"`
	}

	var userReq StockDetails
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&userReq)
	if err != nil {
		http.Error(w, "Bad Request: "+err.Error(), http.StatusBadRequest)
		return
	}

	err = models.BuyStock(userReq.Code, userReq.Cost, userReq.Amount)
	if err != nil {
		http.Error(w, "Internal server error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Stock purchased successfully"))
}

func SellStockHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method should be POST", http.StatusMethodNotAllowed)
		return
	}

	type StockDetails struct {
		Code   string  `json:"code"`
		Cost   float64 `json:"cost"`
		Amount float64 `json:"stock_count"`
	}

	var userReq StockDetails
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&userReq)
	if err != nil {
		http.Error(w, "Bad Request: "+err.Error(), http.StatusBadRequest)
		return
	}

	err = models.SellStock(userReq.Code, userReq.Cost, userReq.Amount)
	if err != nil {
		http.Error(w, "Internal server error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Stock sold successfully"))
}
