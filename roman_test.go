package main

import "testing"

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
