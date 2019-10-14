// Package grains provides functionality to calculate the number of grains of wheat on a chessboard,
// given that the number on each square doubles.
package grains

import (
	"fmt"
	"math"
)

// Square calculates the number of grains on given square of a chessboard.
func Square(v int) (uint64, error) {
	if v <= 0 || v > 64 {
		return 0, fmt.Errorf("failed to calculate for input value %d, should be 0 < input < 65", v)
	}
	return uint64(math.Pow(2, float64(v-1))), nil
}

// Total calculates the total number of grains on the chessboard.
func Total() uint64 {
	var total uint64
	for i := 1; i < 65; i++ {
		square, err := Square(i)
		if err != nil {
			panic(err)
		}
		total += square
	}
	return total
}
