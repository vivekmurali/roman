package roman

import (
	"fmt"
	"testing"
)

func TestRoman(t *testing.T) {
	var tests = []struct {
		name     string
		given    int
		expected string
	}{
		{"Test 1", 100, "C"},
		{"Test 2", 3250, "MMMCCL"},
		{"Test 3", 2199, "MMCXCIX"},
		{"Test 4", 175, "CLXXV"},
		{"Test 5", 1389, "MCCCLXXXIX"},
		{"Test 6", 521, "DXXI"},
		{"Test 7", 123, "CXXIII"},
		{"Test 8", 1998, "MCMXCVIII"},
		{"Test 9", 2000, "MM"},
		{"Test 10", 2021, "MMXXI"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := Roman(tt.given)
			if actual != tt.expected {
				t.Errorf("(%d): expected %s, actual %s", tt.given, tt.expected, actual)
			}

		})
	}
}

func ExampleRoman() {
	s := Roman(100)
	fmt.Println(s)
	// Output: C
}
