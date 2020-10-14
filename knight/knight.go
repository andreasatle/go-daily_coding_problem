// Package knight contains Daily Coding Problem #64
//
// This problem was asked by Google.
//
// A knight's tour is a sequence of moves by a knight on a chessboard
// such that all squares are visited once.
// Given N, write a function to return the number of knight's tours
// on an N by N chessboard.
//
// Note: The board will be extended by two squares in all directions,
// in order to avoid checking out of board.
//
// Author: Andreas Atle, atle.andreas@gmail.com
package knight

// Tuple contains a (row,col)-tuple
type Tuple struct {
	row int
	col int
}

// Contains the different possible knight-moves
var offsets = []Tuple{
	Tuple{-2, -1},
	Tuple{-2, 1},
	Tuple{2, -1},
	Tuple{2, 1},
	Tuple{-1, -2},
	Tuple{-1, 2},
	Tuple{1, -2},
	Tuple{1, 2},
}

// Board contains the moves on a
type Board struct {
	N      int
	N4     int
	nBoard int
	board  []bool
}

// NewBoard creates a new instance of a Board
func NewBoard(n int) *Board {
	b := &Board{N: n, N4: n + 4, nBoard: (n + 4) * (n + 4)}
	b.board = make([]bool, b.nBoard)

	for _, i := range []int{0, 1, b.N4 - 2, b.N4 - 1} {
		for j := 0; j < b.N4; j++ {
			b.board[b.Pos(i, j)] = true
			b.board[b.Pos(j, i)] = true
		}
	}
	return b
}

// Pos returns the position in the slice for a given (row,col)-pair
func (b *Board) Pos(row, col int) int {
	return row*b.N4 + col
}

// Row returns the (extended) row for a given position in the slice
func (b *Board) Row(pos int) int {
	return pos / b.N4
}

// Col returns the (extended) column for a given position in the slice
func (b *Board) Col(pos int) int {
	return pos % b.N4
}

// Filled checks if all positions have been traversed
func (b *Board) Filled() bool {
	for pos := 0; pos < b.nBoard; pos++ {
		if !b.board[pos] {
			return false
		}
	}
	return true
}

// CountToursStartingAt counts all possible tours starting at (row,col)
func (b *Board) CountToursStartingAt(row, col int) int {
	// return 0 if position on board is already set
	pos := b.Pos(row, col)
	if b.board[pos] {
		return 0
	}

	// Set position on board
	b.board[pos] = true

	// Check if all positions have been set
	if b.Filled() {
		return 1
	}

	counter := 0
	oldBoard := make([]bool, len(b.board))
	copy(oldBoard, b.board)

	for _, o := range offsets {
		if !b.board[b.Pos(row+o.row, col+o.col)] {
			counter += b.CountToursStartingAt(row+o.row, col+o.col)
			copy(b.board, oldBoard)
		}
	}
	return counter
}

// CountAllTours counts all tours for all starting positions.
func (b *Board) CountAllTours() int {
	counter := 0
	oldBoard := make([]bool, len(b.board))
	copy(oldBoard, b.board)
	for pos := 0; pos < b.nBoard; pos++ {
		counter += b.CountToursStartingAt(b.Row(pos), b.Col(pos))
		copy(b.board, oldBoard)
	}
	return counter
}
