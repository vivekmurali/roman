package roman

import (
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {
	var tests = []struct {
		name     string
		expected int
		given    string
	}{
		{"Test 1", 100, "C"},
		{"Test 2", 3250, "MMMCCL"},
		{"Test 3", 2199, "MMCXCIX"},
		{"Test 4", 175, "CLXXV"},
		{"Test 5", 1389, "MCCCLXXXIX"},
		{"Test 6", 521, "DxXI"},
		{"Test 7", 123, "cxxiii"},
		{"Test 8", 1998, "MCMXCViii"},
		{"Test 9", 2000, "MM"},
		{"Test 10", 2021, "mmXXI"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := Parse(tt.given)
			if actual != tt.expected {
				t.Errorf("(%s): expected %d, actual %d", tt.given, tt.expected, actual)
			}

		})
	}
}

func ExampleParse() {
	dec := Parse("XII")
	fmt.Println(dec)
	// Output: 12
}
