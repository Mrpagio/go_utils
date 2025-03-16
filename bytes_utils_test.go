package utils

import (
	"fmt"
	"testing"
)

func TestSanitizeUint64Array(t *testing.T) {
	tests := []struct {
		name     string
		input    []uint64
		startIdx int
		length   int
		expected []uint64
	}{
		{
			name:     "Empty array",
			input:    []uint64{},
			startIdx: 0,
			length:   0,
			expected: []uint64{},
		},
		{
			name:     "Array with no sanitization needed",
			input:    []uint64{1, 2, 3, 4, 5},
			startIdx: 0,
			length:   5,
			expected: []uint64{1, 2, 3, 4, 5},
		},
		{
			name:     "Sanitize with startIdx and length",
			input:    []uint64{1, 2, 3, 4, 5, 6, 7, 8},
			startIdx: 2,
			length:   4,
			expected: []uint64{3, 4, 5, 6},
		},
		{
			name:     "Sanitize with startIdx and length exceeding array",
			input:    []uint64{1, 2, 3, 4, 5},
			startIdx: 1,
			length:   10,
			expected: []uint64{2, 3, 4, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SanitizeUint64Array(tt.input, tt.startIdx, tt.length)
			if !equal(result, tt.expected) {
				t.Errorf("SanitizeUint64Array(%v, %d, %d) = %v; want %v", tt.input, tt.startIdx, tt.length, result, tt.expected)
			} else {
				fmt.Printf("Test %s passed: got %v\n", tt.name, result)
			}
		})
	}
}

// Helper function to compare two slices
func equal(a, b []uint64) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
