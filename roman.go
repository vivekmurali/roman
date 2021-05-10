package roman

// The roman function takes in a decimal number and returns the roman numeral as a string
func Roman(n int) string {
	num := []int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}
	str := []string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C", "CD", "D", "CM", "M"}

	i := 12

	var s string

	for n > 0 {
		div := n / num[i]
		n = n % num[i]
		for ; div > 0; div-- {
			s += str[i]
		}
		i--
	}
	return s
}
