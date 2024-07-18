package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestReverseString(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello", "olleh"},
		{"world", "dlrow"},
		{"", ""},
		{"Go", "oG"},
	}

	for _, test := range tests {
		result := ReverseString(test.input)
		if result != test.expected {
			t.Errorf("ReverseString(%q) = %q; want %q", test.input, result, test.expected)
		}
	}
}

func TestReverseHandler(t *testing.T) {
	tests := []struct {
		query    string
		expected string
		status   int
	}{
		{"str=hello", "Reversed string: olleh\n", http.StatusOK},
		{"str=world", "Reversed string: dlrow\n", http.StatusOK},
		{"str=", "No string provided\n", http.StatusBadRequest},
		{"", "No string provided\n", http.StatusBadRequest},
	}

	for _, test := range tests {
		req, err := http.NewRequest("GET", "/reverse?"+test.query, nil)
		if err != nil {
			t.Fatalf("Could not create request: %v", err)
		}
		rec := httptest.NewRecorder()
		reverseHandler(rec, req)

		if rec.Code != test.status {
			t.Errorf("Expected status %d; got %d", test.status, rec.Code)
		}
		if rec.Body.String() != test.expected {
			t.Errorf("Expected body %q; got %q", test.expected, rec.Body.String())
		}
	}
}
