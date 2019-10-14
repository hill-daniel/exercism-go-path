// Package diffsquares provides functionality for finding the difference between the square of the sum and the sum of the squares of the first N natural numbers.
// See also problem 6 at Project Euler http://projecteuler.net/problem=6
package diffsquares

// SquareOfSum calculates square of the sum of a sequence of natural numbers.
// E.g. n = 5 -> (1 + 2 + 3 + 4 + 5)²
func SquareOfSum(n int) int {
	sum := ((n + 1) * n) / 2
	return sum * sum
}

// SumOfSquares calculates the square of each number of a sequence of natural numbers and sums up the results.
// E.g. n = 5 -> 1² + 2² + 3² + 4² + 5²
func SumOfSquares(n int) int {
	var squareSum int
	for i := 1; i <= n; i++ {
		squareSum += i * i
	}
	return squareSum
}

// Difference calculates the difference between SquareOfSum and SumOfSquares for the same n.
// E.g. n = 5 -> (1 + 2 + 3 + 4 + 5)² - (5 -> 1² + 2² + 3² + 4² + 5²)
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
