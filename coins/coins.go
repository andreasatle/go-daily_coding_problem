// Package coins contains Daily Coding Problem #122
//
// This question was asked by Zillow.
//
// You are given a 2-d matrix where each cell represents number of coins in
// that cell. Assuming we start at matrix[0][0], and can only move right or
// down, find the maximum number of coins you can collect by the bottom
// right corner.
//
// For example, in this matrix
//
//  0 3 1 1
//  2 0 0 4
//  1 5 3 1
// The most we can collect is 0 + 2 + 1 + 5 + 3 + 1 = 12 coins.
//
// Author: Andreas Atle, atle.andreas@gmail.com
package coins

// Max returns the maximum coin sum in a matrix traversal
func Max(vals [][]int) int {
	// Special case for empty matrix, avoids boundary checks later
	if len(vals) == 0 || len(vals[0]) == 0 {
		return 0
	}

	acc := createMatrix(vals)
	rows := len(vals)
	cols := len(vals[0])

	acc[rows-1][cols-1] = vals[rows-1][cols-1]
	for col := cols - 2; col >= 0; col-- {
		acc[rows-1][col] = vals[rows-1][col] + acc[rows-1][col+1]
	}
	for row := rows - 2; row >= 0; row-- {
		acc[row][cols-1] = vals[row][cols-1] + acc[row+1][cols-1]
	}

	for row := rows - 2; row >= 0; row-- {
		for col := cols - 2; col >= 0; col-- {
			acc[row][col] = vals[row][col] + max(acc[row+1][col], acc[row][col+1])
		}
	}
	return acc[0][0]
}

func createMatrix(vals [][]int) [][]int {
	mat := make([][]int, len(vals))
	for i := 0; i < len(vals); i++ {
		mat[i] = make([]int, len(vals[i]))
	}
	return mat
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
