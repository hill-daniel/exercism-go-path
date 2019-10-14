// Package isogram provides functionality to handle isograms.
// An isogram (also known as a "nonpattern word") is a word or phrase without a repeating letter
// however spaces and hyphens are allowed to appear multiple times.
package isogram

import (
	"unicode"
)

// IsIsogram determines if a given string is an isogram,
// e.g. a word or phrase without a repeating letter, ignoring spaces and hyphens.
func IsIsogram(word string) bool {
	seenRunes := make(map[rune]bool)
	for _, r := range word {
		if !unicode.IsLetter(r) {
			continue
		}
		upperRune := unicode.ToUpper(r)
		if seenRunes[upperRune] {
			return false
		}
		seenRunes[upperRune] = true
	}
	return true
}
