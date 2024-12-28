package models

import (
    "database/sql"
)

type User struct {
    ID       int
    Username string
    Password string
    Balance  float64
}

func CreateUser(db *sql.DB, username, password string) error {
    query := `INSERT INTO users (username, password_hash) VALUES ($1, $2)`
    _, err := db.Exec(query, username, password)
    return err
}

func GetUser(db *sql.DB, username string) (*User, error) {
    var user User
    query := `SELECT id, username, password_hash, balance FROM users WHERE username=$1`
    err := db.QueryRow(query, username).Scan(&user.ID, &user.Username, &user.Password, &user.Balance)
    if err != nil {
        return nil, err
    }
    return &user, nil
}
