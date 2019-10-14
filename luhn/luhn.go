// Package luhn provides functionality to check validity of a given number.
// The Luhn algorithm is a simple checksum formula used to validate a variety of identification
// numbers, such as credit card numbers and Canadian Social Insurance Numbers.
package luhn

import (
	"strings"
	"unicode"
)

// Valid checks validity of a given input via checksum calculation.
// Returns true if valid, false if not.
func Valid(input string) bool {
	input = strings.Replace(input, " ", "", -1)
	if len(input) <= 1 {
		return false
	}
	var sum int
	for i, r := range input {
		if !unicode.IsNumber(r) {
			return false
		}
		value := int(r - '0')
		if (len(input)+i)%2 == 0 {
			value *= 2
			if value > 9 {
				value -= 9
			}
		}
		sum += value
	}
	if sum%10 == 0 {
		return true
	}
	return false
}

// Valid2 determines whether the digits in the given string constitute a valid luhn code,
// but performance optimized.
func Valid2(s string) bool {
	var n, d, i, m int
	for i = len(s) - 1; i >= 0; i-- {
		c := s[i]
		switch {
		case c == ' ':
			continue
		case c >= '0' && c <= '9':
			m = int(c - '0')
			if d%2 == 1 {
				m <<= 1
			}
			if m > 9 {
				m -= 9
			}
			n += m
			d++
		default:
			return false
		}
	}
	return d > 1 && n%10 == 0
}