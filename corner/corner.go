// Package corner contains Daily Coding Problem #62
//
// This problem was asked by Facebook.
//
// There is an N by M matrix of zeroes. Given N and M, write a function to
// count the number of ways of starting at the top-left corner and getting
// to the bottom-right corner. You can only move right or down.
//
// For example, given a 2 by 2 matrix, you should return 2, since there are
// two ways to get to the bottom-right:
//
//  Right, then down
//  Down, then right
//
// Given a 5 by 5 matrix, there are 70 ways to get to the bottom-right.
//
// Note: The problem is solved in two ways. CountPathsV1 solves the problem
// with a simple recursion. CountPathsV2 solves the problem with counting
// the cases backwards.
//
// Author: Andreas Atle, atle.andreas@gmail.com
package corner

// CountPathsV1 returns the number of paths using a recursive algorithm.
func CountPathsV1(rows, cols int) int {
	// No path
	if rows <= 0 || cols <= 0 {
		return 0
	}
	// Can only go in one direction
	if rows <= 1 || cols <= 1 {
		return 1
	}

	return CountPathsV1(rows-1, cols) + CountPathsV1(rows, cols-1)
}

// CountPathsV1 returns the number of paths using a "dynamic programming"
// approach.
func CountPathsV2(rows, cols int) int {
	if rows <= 0 || cols <= 0 {
		return 0
	}

	// Allocate matrix
	mat := make([][]int, rows)
	for row := 0; row < rows; row++ {
		mat[row] = make([]int, cols)
	}

	// Initialize matrix at boundary
	for col := 0; col < cols; col++ {
		mat[0][col] = 1
	}
	for row := 0; row < rows; row++ {
		mat[row][0] = 1
	}

	// Fill in matrix
	for row := 1; row < rows; row++ {
		for col := 1; col < cols; col++ {
			mat[row][col] = mat[row-1][col] + mat[row][col-1]
		}
	}
	return mat[rows-1][cols-1]
}
