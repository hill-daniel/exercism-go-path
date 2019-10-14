// Package scrabble offers functionality for the scrabble game.
package scrabble

import "unicode"

var letterToScore = map[string]int{
	"A": 1,
	"E": 1,
	"I": 1,
	"O": 1,
	"U": 1,
	"L": 1,
	"N": 1,
	"R": 1,
	"S": 1,
	"T": 1,
	"D": 2,
	"G": 2,
	"B": 3,
	"C": 3,
	"M": 3,
	"P": 3,
	"F": 4,
	"H": 4,
	"V": 4,
	"W": 4,
	"Y": 4,
	"K": 5,
	"J": 8,
	"X": 8,
	"Q": 10,
	"Z": 10,
}

// Score computes the scrabble score for a given word.
// E.g.:
//	Letter                           Value
//	A, E, I, O, U, L, N, R, S, T       1
//	D, G                               2
//	B, C, M, P                         3
//	F, H, V, W, Y                      4
//	K                                  5
//	J, X                               8
//	Q, Z                               10
func Score(word string) int {
	var score int
	for _, r := range word {
		letter := string(unicode.ToUpper(r))
		if letterScore, ok := letterToScore[letter]; ok {
			score += letterScore
		}
	}
	return score
}

// Further improvement
// you could use a map[rune]int instead of map[string]int for direct lookup of the score value
// for a rune: no type conversion needed. A rune is created with single quotes e.g. 'A' just like
// a byte. if you are up for it using a switch instead of a map will increase speed significantly
var runeToScore = map[rune]int{
	'A': 1,
	'E': 1,
	'I': 1,
	'O': 1,
	'U': 1,
	'L': 1,
	'N': 1,
	'R': 1,
	'S': 1,
	'T': 1,
	'D': 2,
	'G': 2,
	'B': 3,
	'C': 3,
	'M': 3,
	'P': 3,
	'F': 4,
	'H': 4,
	'V': 4,
	'W': 4,
	'Y': 4,
	'K': 5,
	'J': 8,
	'X': 8,
	'Q': 10,
	'Z': 10,
}

// Score2 does the same as Score but with a rune map.
// Faster than Score, but slower than Score3
func Score2(word string) int {
	var score int
	for _, r := range word {
		if runeScore, ok := runeToScore[unicode.ToUpper(r)]; ok {
			score += runeScore
		}
	}
	return score
}

// Score3 does the same as Score but with a switch statement
// Note: Fastest impl. of the three.
func Score3(word string) int {
	var score int
	for _, r := range word {
		switch unicode.ToUpper(r) {
		case 'A':
			fallthrough
		case 'E':
			fallthrough
		case 'I':
			fallthrough
		case 'O':
			fallthrough
		case 'U':
			fallthrough
		case 'L':
			fallthrough
		case 'N':
			fallthrough
		case 'R':
			fallthrough
		case 'S':
			fallthrough
		case 'T':
			score++
		case 'D':
			fallthrough
		case 'G':
			score += 2
		case 'B':
			fallthrough
		case 'C':
			fallthrough
		case 'M':
			fallthrough
		case 'P':
			score += 3
		case 'F':
			fallthrough
		case 'H':
			fallthrough
		case 'V':
			fallthrough
		case 'W':
			fallthrough
		case 'Y':
			score += 4
		case 'K':
			score += 5
		case 'J':
			fallthrough
		case 'X':
			score += 8
		case 'Q':
			fallthrough
		case 'Z':
			score += 10
		}
	}
	return score
}
