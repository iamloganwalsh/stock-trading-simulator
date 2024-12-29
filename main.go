package main

import (

	"log"
	"net/http"

	"github.com/iamloganwalsh/stock-trading-simulator/routes"
	"github.com/iamloganwalsh/stock-trading-simulator/config"
)



func main() {

	db, err := config.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	err = config.InitDB(db)
	if err != nil {
		log.Fatalf("Failed to initialize the database: %v", err)
	}
	

	http.HandleFunc("/user/register", routes.RegisterUser)
	http.HandleFunc("/user/login", routes.LoginUser)

	log.Println("Starting server on localhost:3000...")
	log.Fatal(http.ListenAndServe("localhost:3000", nil))
}
