// Package matrix provides functionality for handling matrices
package matrix

import (
	"github.com/pkg/errors"
	"strconv"
	"strings"
)

// Matrix represents a two dimensional matrix
type Matrix struct {
	table [][]int
}

// New creates a new Matrix from a given string,
// such as that "1 4 9\n16 25 36" will create a matrix representing:
//      0   1   2
//   |-----------
// 0 |  1   4   9
// 1 | 16  15  36
func New(input string) (*Matrix, error) {
	table, err := createTable(input)
	if err != nil {
		return nil, err
	}
	return &Matrix{table: table}, nil
}

// Set sets a given int value on the given row and column index
func (m Matrix) Set(row int, col int, val int) bool {
	if row > len(m.table)-1 || row < 0 || col > len(m.table[0])-1 || col < 0 {
		return false
	}
	m.table[row][col] = val
	return true
}

// Rows returns the rows of the matrix e.g.
//     0   1   2
//   |-----------
// 0 |  1   4   9
// 1 | 16  15  36
// returns
// 1   4   9
// 16  15  36
func (m Matrix) Rows() [][]int {
	return copyTable(m.table)
}

// Cols returns the columns of the matrix e.g.
//     0   1   2
//   |-----------
// 0 |  1   4   9
// 1 | 16  15  36
// returns
// 1   16
// 4   15
// 9   36
func (m Matrix) Cols() [][]int {
	return transpose(m.table)
}

func createTable(input string) ([][]int, error) {
	rows := strings.Split(input, "\n")
	table := make([][]int, len(rows))
	for i, row := range rows {
		columns, err := createColumns(strings.Fields(row))
		if err != nil {
			return nil, err
		}
		if i > 0 && len(columns) != len(table[0]) {
			return nil, errors.New("uneven length of row")
		}
		table[i] = columns
	}
	return table, nil
}

func createColumns(fields []string) ([]int, error) {
	columns := make([]int, len(fields))
	for j, field := range fields {
		value, err := strconv.Atoi(field)
		if err != nil {
			return nil, err
		}
		columns[j] = value
	}
	return columns, nil
}

func copyTable(table [][]int) [][]int {
	tableCopy := make([][]int, len(table))
	for i, row := range table {
		rowCopy := make([]int, len(row))
		copy(rowCopy, row)
		tableCopy[i] = rowCopy
	}
	return tableCopy
}

func transpose(table [][]int) [][]int {
	rowCount := len(table)
	colCount := len(table[0])
	tableCopy := make([][]int, colCount)
	for i := 0; i < colCount; i++ {
		transCol := make([]int, rowCount)
		for j := 0; j < rowCount; j++ {
			transCol[j] = table[j][i]
		}
		tableCopy[i] = transCol
	}
	return tableCopy
}
