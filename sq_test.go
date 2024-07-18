package main

import (
	"testing"
)

func TestSquare(t *testing.T) {
	tests := []struct {
		a        float64
		expected float64
	}{
		{3, 9},
		{-3, 9},
		{0, 0},
		{1.5, 2.25},
	}

	for _, test := range tests {
		result := Square(test.a)
		if result != test.expected {
			t.Errorf("Square(%v) = %v; want %v", test.a, result, test.expected)
		}
	}
}
