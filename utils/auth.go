package utils

import (
    "golang.org/x/crypto/bcrypt"
    "log"
)

func HashPassword(password string) string {
    hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        log.Fatalf("Error hashing password: %v", err)
    }
    return string(hash)
}

func VerifyPassword(hashedPassword, password string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
    return err == nil
}
