// Package letter provides functionality for processing text.
package letter

// FreqMap stores the count of occurrence of a rune
type FreqMap map[rune]int

// Frequency calculates count of occurrence of a letter in given string.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency calculates count of occurrence of a letter in given string, but concurrently.
func ConcurrentFrequency(ss []string) FreqMap {
	result := FreqMap{}
	fc := make(chan FreqMap)
	for _, s := range ss {
		go func(value string) {
			fc <- Frequency(value)
		}(s)
	}
	for range ss {
		for k, v := range <-fc {
			result[k] += v
		}
	}
	return result
}
