package main

import (

	"log"
	"net/http"
	"github.com/iamloganwalsh/stock-trading-simulator/config"
)



func main() {

	// Making sure database exists & if not then create it
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	err = config.InitDB(db)
	if err != nil {
		log.Fatalf("Failed to initialize the database: %v", err)
	}


	log.Println("Starting server on localhost:3000...")
	log.Fatal(http.ListenAndServe("localhost:3000", nil))
}
