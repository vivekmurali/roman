package main

import (
	"strings"
)

func Parse(s string) int {
	val := make(map[string]int)
	val["I"] = 1
	val["V"] = 5
	val["X"] = 10
	val["L"] = 50
	val["C"] = 100
	val["D"] = 500
	val["M"] = 1000

	letters := strings.Split(s, "")

	sum := 0
	for i := 0; i < len(letters); i++ {
		if i == len(letters)-1 {
			sum += val[letters[i]]
			break
		}
		if val[letters[i]] < val[letters[i+1]] {
			sum += val[letters[i+1]] - val[letters[i]]
			i++
			continue
		}
		sum += val[letters[i]]
	}
	return sum
}
