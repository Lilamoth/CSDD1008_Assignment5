package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGeneratePassword(t *testing.T) {
	tests := []struct {
		length   int
		contains bool
	}{
		{8, true},
		{12, true},
		{16, true},
		{32, true},
		{0, false},
		{-5, false},
	}

	for _, test := range tests {
		password, err := generatePassword(test.length)
		if (err == nil) != test.contains {
			t.Errorf("generatePassword(%d) error: %v; want contains %v", test.length, err, test.contains)
		}
		if test.contains && len(password) != test.length {
			t.Errorf("generatePassword(%d) returned password of length %d; want %d", test.length, len(password), test.length)
		}
	}
}

func TestPasswordHandler(t *testing.T) {
	tests := []struct {
		query    string
		expected string
		status   int
	}{
		{"length=8", "Generated password: ", http.StatusOK},
		{"length=12", "Generated password: ", http.StatusOK},
		{"length=0", "Invalid input; length must be a positive integer\n", http.StatusBadRequest},
		{"length=-5", "Invalid input; length must be a positive integer\n", http.StatusBadRequest},
		{"length=abc", "Invalid input; length must be a positive integer\n", http.StatusBadRequest},
	}

	for _, test := range tests {
		req, err := http.NewRequest("GET", "/password?"+test.query, nil)
		if err != nil {
			t.Fatalf("Could not create request: %v", err)
		}
		rec := httptest.NewRecorder()
		passwordHandler(rec, req)

		if rec.Code != test.status {
			t.Errorf("Expected status %d; got %d", test.status, rec.Code)
		}
		if test.status == http.StatusOK {
			if !strings.HasPrefix(rec.Body.String(), test.expected) {
				t.Errorf("Expected body starting with %q; got %q", test.expected, rec.Body.String())
			}
		} else {
			if rec.Body.String() != test.expected {
				t.Errorf("Expected body %q; got %q", test.expected, rec.Body.String())
			}
		}
	}
}
