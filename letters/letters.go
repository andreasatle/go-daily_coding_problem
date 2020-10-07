// Package letters contains Daily Coding Problem #98
//
//This problem was asked by Coursera.
//
// Given a 2D board of characters and a word, find if the word exists in the grid.
//
// The word can be constructed from letters of sequentially adjacent cell,
// where "adjacent" cells are those horizontally or vertically neighboring.
// The same letter cell may not be used more than once.
//
// For example, given the following board:
//  [
//    ['A','B','C','E'],
//    ['S','F','C','S'],
//    ['A','D','E','E']
//  ]
// exists(board, "ABCCED") returns true,
// exists(board, "SEE") returns true,
// exists(board, "ABCB") returns false.
//
// Note: The problem is solved using back-tracking. We also use defer to
// simplify the cleaning up of the mask after use in recursion.
//
// Author: Andreas Atle, atle.andreas@gmail.com
package letters

import (
	"errors"
)

// Board implements a representaion of a board
type Board struct {
	Rows int
	Cols int
	Strs []string
}

// NewBoard returns a new Board-pointer
func NewBoard(strs []string) (*Board, error) {

	if len(strs) == 0 {
		return nil, errors.New("empty board")
	}

	b := &Board{}
	b.Rows = len(strs)
	b.Cols = len(strs[0])

	for _, str := range strs {
		if b.Cols != len(str) {
			return nil, errors.New("string size mismatch")
		}
		b.Strs = append(b.Strs, str)
	}

	return b, nil
}

// NewMask returns a new mask containing [][]bool
func (b *Board) NewMask() [][]bool {

	mask := make([][]bool, b.Rows)

	for i := range mask {
		mask[i] = make([]bool, b.Cols)
	}

	return mask
}

// PathExists returns true if a path defined by str exists without cycles
func (b *Board) PathExists(str string) bool {

	if len(str) == 0 {
		return true
	}

	mask := b.NewMask()
	for row := 0; row < b.Rows; row++ {
		for col := 0; col < b.Cols; col++ {
			if str[0] == b.Strs[row][col] {
				if b.pathExists(str, mask, row, col) {
					return true
				}
			}
		}
	}

	return false
}

func (b *Board) pathExists(str string, mask [][]bool, row, col int) bool {

	if len(str) == 0 {
		return true
	}

	if row < 0 || col < 0 || row >= b.Rows || col >= b.Cols {
		return false
	}

	if !mask[row][col] && b.Strs[row][col] == str[0] {
		mask[row][col] = true
		defer func() { mask[row][col] = false }()

		if b.pathExists(str[1:], mask, row-1, col) {
			return true
		}
		if b.pathExists(str[1:], mask, row+1, col) {
			return true
		}
		if b.pathExists(str[1:], mask, row, col-1) {
			return true
		}
		if b.pathExists(str[1:], mask, row, col+1) {
			return true
		}
	}

	return false
}
