// bob provides functionality to simulate interaction with the lackadaisical teenager bob.
package bob

import (
	"strings"
	"unicode"
)

// Returns a reaction from bob, depending on the style of the provided string input.
func Hey(remark string) string {
	remark = removeWhitespace(remark)
	if isEmpty(remark) {
		return "Fine. Be that way!"
	}
	if isQuestion(remark) {
		if isShouted(remark) && containsLetter(remark) {
			return "Calm down, I know what I'm doing!"
		}
		return "Sure."
	}
	if isShouted(remark) && containsLetter(remark) {
		return "Whoa, chill out!"
	}
	return "Whatever."
}

func removeWhitespace(v string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, v)
}

func isEmpty(remark string) bool {
	return len(remark) == 0
}

func isQuestion(remark string) bool {
	return strings.HasSuffix(remark, "?")
}

func isShouted(remark string) bool {
	return strings.ToUpper(remark) == remark
}

func containsLetter(s string) bool {
	for _, r := range s {
		if unicode.IsLetter(r) {
			return true
		}
	}
	return false
}
