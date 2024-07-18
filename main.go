package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"strconv"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const numberBytes = "0123456789"
const specialBytes = "!@#$%^&*()-_=+[]{}|;:,.<>?/~"

func generatePassword(length int) (string, error) {
	if length <= 0 {
		return "", fmt.Errorf("length must be a positive integer")
	}

	allChars := letterBytes + numberBytes + specialBytes
	password := make([]byte, length)
	for i := range password {
		idx, err := rand.Int(rand.Reader, big.NewInt(int64(len(allChars))))
		if err != nil {
			return "", err
		}
		password[i] = allChars[idx.Int64()]
	}
	return string(password), nil
}

// passwordHandler handles requests to the "/password" URL
func passwordHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	lengthStr := query.Get("length")

	length, err := strconv.Atoi(lengthStr)
	if err != nil || length <= 0 {
		http.Error(w, "Invalid input; length must be a positive integer", http.StatusBadRequest)
		return
	}

	password, err := generatePassword(length)
	if err != nil {
		http.Error(w, "Error generating password", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Generated password: %s\n", password)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/password", passwordHandler)

	log.Println("Starting server on :3001")
	err := http.ListenAndServe(":3001", mux)
	if err != nil {
		log.Fatalf("Error starting server: %s\n", err)
	}
}
