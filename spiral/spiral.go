// Package spiral contains Daily Coding Problem #65
//
// This problem was asked by Amazon.
//
// Given a N by M matrix of numbers, print out the matrix in a clockwise spiral.
//
// For example, given the following matrix:
//
//  [[1,  2,  3,  4,  5],
//   [6,  7,  8,  9,  10],
//   [11, 12, 13, 14, 15],
//   [16, 17, 18, 19, 20]]
//
// You should print out the following:
//
//  1
//  2
//  3
//  4
//  5
//  10
//  15
//  20
//  19
//  18
//  17
//  16
//  11
//  6
//  7
//  8
//  9
//  14
//  13
//  12
//
// Author: Andreas Atle, atle.andreas@gmail.com
package spiral

import (
	"errors"
)

// ToSlice return matrix elements with clockwise spiral ordering
func ToSlice(mat [][]int) ([]int, error) {
	// Check that the matrix is rectangular
	rows, cols, err := validate(mat)
	if err != nil {
		return []int{}, err
	}

	minRow, minCol, maxRow, maxCol := 0, 0, rows-1, cols-1
	output := make([]int, rows*cols)
	cnt := 0
	for minRow <= maxRow && minCol <= maxCol {

		// Sweep right at upper edge
		for col := minCol; col <= maxCol; col++ {
			output[cnt] = mat[minRow][col]
			cnt++
		}
		minRow++

		// Sweep down at right edge
		for row := minRow; row <= maxRow; row++ {
			output[cnt] = mat[row][maxCol]
			cnt++
		}
		maxCol--

		// Sweep left at bottom edge
		for col := maxCol; col >= minCol; col-- {
			output[cnt] = mat[maxRow][col]
			cnt++
		}
		maxRow--

		// Sweep up at left edge
		for row := maxRow; row >= minRow; row-- {
			output[cnt] = mat[row][minCol]
			cnt++
		}
		minCol++
	}
	return output, err
}

func validate(mat [][]int) (int, int, error) {

	rows := len(mat)

	if rows == 0 {
		return 0, 0, nil
	}

	cols := len(mat[0])
	for _, row := range mat {
		if len(row) != cols {
			return -1, -1, errors.New("non-rectangular matrix")
		}
	}

	return rows, cols, nil
}
