package routes

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/iamloganwalsh/stock-trading-simulator/utils"
)

func FetchCryptoHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Method should be GET", http.StatusMethodNotAllowed)
		return
	}

	vars := mux.Vars(r)
	code := vars["code"]

	price, err := utils.Fetch_crypto_price(code)
	if err != nil {
		http.Error(w, "Internal server error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	crypto_price := fmt.Sprintf("%f", price)

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte(crypto_price)); err != nil {
		http.Error(w, "Failed to write response: "+err.Error(), http.StatusInternalServerError)
	}
}

func FetchStockHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Method should be GET", http.StatusMethodNotAllowed)
		return
	}

	vars := mux.Vars(r)
	code := vars["code"]

	price, err := utils.Fetch_crypto_price(code)
	if err != nil {
		http.Error(w, "Internal server error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	crypto_price := fmt.Sprintf("%f", price)

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte(crypto_price)); err != nil {
		http.Error(w, "Failed to write response: "+err.Error(), http.StatusInternalServerError)
	}
}

func FetchCryptoPrevHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Method should be GET", http.StatusMethodNotAllowed)
		return
	}

	vars := mux.Vars(r)
	code := vars["code"]

	price, err := utils.Fetch_prev_crypto(code)
	if err != nil {
		http.Error(w, "Internal server error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	crypto_price := fmt.Sprintf("%v", price)

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte(crypto_price)); err != nil {
		http.Error(w, "Failed to write response: "+err.Error(), http.StatusInternalServerError)
	}
}
