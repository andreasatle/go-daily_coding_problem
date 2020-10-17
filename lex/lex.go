// Package lex contains Daily Coding Problem #76
//
// This problem was asked by Google.
//
// You are given an N by M 2D matrix of lowercase letters. Determine the
// minimum number of columns that can be removed to ensure that each row
// is ordered from top to bottom lexicographically. That is, the letter at
// each column is lexicographically later as you go down each row. It does
// not matter whether each row itself is ordered lexicographically.
//
// For example, given the following table:
//
//  cba
//  daf
//  ghi
// This is not ordered because of the a in the center. We can remove the
// second column to make it ordered:
//
//  ca
//  df
//  gi
// So your function should return 1, since we only needed to remove 1 column.
//
// As another example, given the following table:
//
//  abcdef
// Your function should return 0, since the rows are already ordered
// (there's only one row).
//
// As another example, given the following table:
//
//  zyx
//  wvu
//  tsr
// Your function should return 3, since we would need to remove all the
// columns to order it.
//
// Author: Andreas Atle, atle.andreas@gmail.com
package lex

import (
	"errors"
)

// NonOrderedColumns returns the number of columns that are not
// lexicographically ordered
func NonOrderedColumns(rows []string) (int, error) {
	if len(rows) <= 1 {
		return 0, nil
	}

	if !checkRows(rows) {
		return -1, errors.New("strings are of variable length")
	}

	colCnt := 0
	for col := 0; col < len(rows[0]); col++ {
		for row := 1; row < len(rows); row++ {
			if rows[row][col] < rows[row-1][col] {
				colCnt++
				break
			}
		}
	}
	return colCnt, nil
}

// checkRows returns true if all rows are of the same length
func checkRows(rows []string) bool {
	for _, row := range rows {
		if len(row) != len(rows[0]) {
			return false
		}
	}
	return true
}
