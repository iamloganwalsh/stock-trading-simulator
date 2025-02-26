package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/iamloganwalsh/stock-trading-simulator/config"
	"github.com/iamloganwalsh/stock-trading-simulator/models"
)

func InitUserHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method should be POST", http.StatusMethodNotAllowed)
		return
	}

	type InitUserRequest struct {
		Username string `json:"username"`
	}

	var userReq InitUserRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&userReq)
	if err != nil {
		http.Error(w, "Bad Request: "+err.Error(), http.StatusBadRequest)
		return
	}

	db, _ := config.ConnectDB()

	err = models.InitUser(db, userReq.Username)
	if err != nil {
		http.Error(w, "Internal server error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User created successfully"))
}

func GetUsernameHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Method should be GET", http.StatusMethodNotAllowed)
		return
	}
	db, _ := config.ConnectDB()
	username, err := models.GetUsername(db)
	if err != nil {
		http.Error(w, "Internal server error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte(username)); err != nil {
		http.Error(w, "Failed to write response: "+err.Error(), http.StatusInternalServerError)
	}
}

func GetBalanceHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Method should be GET", http.StatusMethodNotAllowed)
		return
	}
	db, _ := config.ConnectDB()
	balance, err := models.GetBalance(db)
	if err != nil {
		http.Error(w, "Internal server error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	balance_string := fmt.Sprintf("%f", balance)

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte(balance_string)); err != nil {
		http.Error(w, "Failed to write response: "+err.Error(), http.StatusInternalServerError)
	}
}

func GetInvestmentHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Method should be GET", http.StatusMethodNotAllowed)
		return
	}
	db, _ := config.ConnectDB()
	investment, err := models.GetInvestment(db)
	if err != nil {
		http.Error(w, "Internal server error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	investment_string := fmt.Sprintf("%f", investment)

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte(investment_string)); err != nil {
		http.Error(w, "Failed to write response: "+err.Error(), http.StatusInternalServerError)
	}
}

func GetCryptoPortfolioHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Method should be GET", http.StatusMethodNotAllowed)
		return
	}
	db, _ := config.ConnectDB()
	crypto_items, err := models.GetCryptoPortfolio(db)
	if err != nil {
		http.Error(w, "Internal server error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	json_data, err := json.Marshal(crypto_items)
	if err != nil {
		http.Error(w, "Failed to encode data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json_data)
}

func GetStockPortfolioHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Method should be GET", http.StatusMethodNotAllowed)
		return
	}
	db, _ := config.ConnectDB()
	stock_items, err := models.GetStockPortfolio(db)
	if err != nil {
		http.Error(w, "Internal server error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	json_data, err := json.Marshal(stock_items)
	if err != nil {
		http.Error(w, "Failed to encode data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json_data)
}

func GetTradeHistoryHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Method should be GET", http.StatusMethodNotAllowed)
		return
	}
	db, _ := config.ConnectDB()
	stock_items, err := models.GetTradeHistory(db)
	if err != nil {
		http.Error(w, "Internal server error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	json_data, err := json.Marshal(stock_items)
	if err != nil {
		http.Error(w, "Failed to encode data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json_data)
}
