package routes

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"github.com/iamloganwalsh/stock-trading-simulator/models"
	"github.com/iamloganwalsh/stock-trading-simulator/utils"
	"fmt"
)

var db *sql.DB // Initialize this in your main setup or config package

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	// Log the incoming request method and URL
	fmt.Println("Received request:", r.Method, r.URL)

	// Only allow POST requests for registration
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Parse the incoming request body
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// For now, just log the received user data
	fmt.Println("Received user data:", user)

	// Respond with success (you can modify this to store data in DB later)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"status": "success"}`))
}

// LoginUser handles user login
func LoginUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var loginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	json.NewDecoder(r.Body).Decode(&loginRequest)

	// Retrieve user from the database
	user, err := models.GetUser(db, loginRequest.Username)
	if err != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	// Verify password
	if !utils.VerifyPassword(user.Password, loginRequest.Password) {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	// Respond with success
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Login successful"))
}
