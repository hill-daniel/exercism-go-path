// Package collatzconjecture provides functionality to handle CollatzConjecture.
// The Collatz Conjecture or 3x+1 problem can be summarized as follows:
// Take any positive integer n. If n is even, divide n by 2 to get n / 2.
// If n is odd, multiply n by 3 and add 1 to get 3n + 1. Repeat the process indefinitely.
// The conjecture states that no matter which number you start with,
// you will always reach 1 eventually.
package collatzconjecture

import "errors"

// CollatzConjecture returns the number of steps required to reach 1, given a number n.
// The number has to be greater that zero, otherwise -1 and an error is returned.
func CollatzConjecture(n int) (int, error) {
	if n <= 0 {
		return -1, errors.New("failed to calculate CollatzConjecture. Input is zero or negative")
	}
	return calcConjecture(n, 0), nil
}

func calcConjecture(n int, steps int) int {
	if n == 1 {
		return steps
	}
	if n%2 == 0 {
		return calcConjecture(n/2, steps+1)
	}
	return calcConjecture(n*3+1, steps+1)
}
