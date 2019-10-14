// Package hamming provides functionality for handling 'Hamming distance'.
// By counting the number of differences between two homologous DNA strands taken from different
// genomes with a common ancestor, we get a measure of the minimum number of point mutations that
// could have occurred on the evolutionary path between the two strands.
//
// This is called the 'Hamming distance'.
package hamming

import "errors"

// Distance calculates the 'Hamming distance' between two DNA strands.
// E.g. compares two DNA strands (strings) and counts how many of the nucleotides are different from their equivalent in the other string.
// returns -1 and error for strands with unequal length.
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, errors.New("failed to calculate distance. Strands have unequal length")
	}
	var distance int
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			distance++
		}
	}
	return distance, nil
}
