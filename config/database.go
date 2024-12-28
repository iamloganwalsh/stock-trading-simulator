package config

import (
    "database/sql"
    _ "github.com/lib/pq" // PostgreSQL driver
)

func ConnectDB() (*sql.DB, error) {
    connStr := "user=username dbname=stocktrader sslmode=disable"
    return sql.Open("postgres", connStr)
}
