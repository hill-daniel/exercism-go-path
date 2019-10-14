// Package twelve provides functionality to print a christmas song.
package twelve

import (
	"fmt"
	"strings"
)

const lyricTemplate = "On the %s day of Christmas my true love gave to me: %s."

var gifts = []string{"a Partridge in a Pear Tree", "two Turtle Doves", "three French Hens", "four Calling Birds", "five Gold Rings", "six Geese-a-Laying", "seven Swans-a-Swimming", "eight Maids-a-Milking", "nine Ladies Dancing", "ten Lords-a-Leaping", "eleven Pipers Piping", "twelve Drummers Drumming"}
var days = []string{"first", "second", "third", "fourth", "fifth", "sixth", "seventh", "eighth", "ninth", "tenth", "eleventh", "twelfth"}

// Song returns all verses of the "Twelve days of christmas" song as string.
func Song() string {
	builder := strings.Builder{}
	for i := 0; i < len(gifts); i++ {
		builder.WriteString(verse(i))
		builder.WriteString("\n")
	}
	return builder.String()
}

// Verse returns a verses of the "Twelve days of christmas" song as string.
func Verse(line int) string {
	return verse(line - 1)
}

func verse(value int) string {
	builder := strings.Builder{}
	for i := value; i >= 0; i-- {
		if i == value {
			builder.WriteString(gifts[i])
		} else if i == 0 {
			builder.WriteString(", and ")
			builder.WriteString(gifts[i])
		} else {
			builder.WriteString(", ")
			builder.WriteString(gifts[i])
		}
	}
	return fmt.Sprintf(lyricTemplate, days[value], builder.String())
}
