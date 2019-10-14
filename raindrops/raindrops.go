// Package raindrops provides functionality to convert a number to a string.
package raindrops

import "strconv"

// Convert converts a number to a string, the contents of which depend on the number's factors.
// E.g.:
//
// * If the number has 3 as a factor, output 'Pling'.
// * If the number has 5 as a factor, output 'Plang'.
// * If the number has 7 as a factor, output 'Plong'.
// * If the number does not have 3, 5, or 7 as a factor, just pass the number's digits straight through.
func Convert(input int) string {
	var result string
	if input%3 == 0 {
		result += "Pling"
	}
	if input%5 == 0 {
		result += "Plang"
	}
	if input%7 == 0 {
		result += "Plong"
	}
	if len(result) == 0 {
		result = strconv.Itoa(input)
	}
	return result
}
