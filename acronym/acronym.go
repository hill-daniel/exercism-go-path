// Provides functionality to handle acronyms.
package acronym

import (
	"strings"
)

// Given a string value, concatenates the first letter of each word and returns
// them in uppercase. E.g. First in first out -> FIFO.
// Considers hyphen separated input also as individual words. E.g. metal-oxide -> MO.
func Abbreviate(s string) string {
	return strings.ToUpper(acro(s, " "))
}

func acro(s string, sep string) string {
	var acronym string
	for _, word := range strings.Split(s, sep) {
		if strings.Contains(s, "-") {
			acronym += acro(word, "-")
		} else {
			acronym += string(word[0])
		}
	}
	return acronym
}
