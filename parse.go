/*

The roman package provides two functions, parse and roman.
This package is used to change between decimals and roman numerals.

*/
package roman

import (
	"strings"
)

// Parses a roman numeral input and returns it's decimal value
// As of now, there is no validation of the roman numeral
// All non roman characters take a value of 0
func Parse(s string) int {
	val := make(map[string]int)
	val["I"] = 1
	val["V"] = 5
	val["X"] = 10
	val["L"] = 50
	val["C"] = 100
	val["D"] = 500
	val["M"] = 1000

	s = strings.ToUpper(s)
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
