package utils

import "strconv"

// Converts an integer to a string with a specified number of digits,
func NumberToDigit(num int, n int) string {
	// Convert the number to a string
	str := strconv.Itoa(num)

	// If the string is already n characters long, return it as is
	if len(str) == n {
		return str
	}

	// If the string is shorter than n characters, pad it with zeros at the front
	for len(str) < n {
		str = "0" + str
	}

	return str
}
