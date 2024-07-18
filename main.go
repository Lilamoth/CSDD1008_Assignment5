package main

import (
	"fmt"
	"log"
	"net/http"
)

// ReverseString reverses a given string
func ReverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// reverseHandler handles requests to the "/reverse" URL
func reverseHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	str := query.Get("str")

	if str == "" {
		http.Error(w, "No string provided", http.StatusBadRequest)
		return
	}

	reversed := ReverseString(str)
	fmt.Fprintf(w, "Reversed string: %s\n", reversed)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/reverse", reverseHandler)

	log.Println("Starting server on :3001")
	err := http.ListenAndServe(":3001", mux)
	if err != nil {
		log.Fatalf("Error starting server: %s\n", err)
	}
}
