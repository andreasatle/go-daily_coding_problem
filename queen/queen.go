// Package queen contains Daily Coding Problem #38
//
// This problem was asked by Microsoft.
//
// You have an N by N board. Write a function that, given N,
// returns the number of possible arrangements of the board
// where N queens can be placed on the board without threatening
// each other, i.e. no two queens share the same row, column, or
// diagonal.
//
// Note: If there is a solution, then each row (as well as column) will
// have exactly one queen. It is therefore sufficient to add the first
// queen in row 0 (for instance). We could furthermore use symmetry of the
// queens move to reduce the initial position to half ((N+1)/2) positions.
// For simplicity, we will start at each column of row 0.
//
// Note: The solution to this problem: (https://en.wikipedia.org/wiki/Eight_queens_puzzle)
//  ---------------------------------
//  N          1  2  3  4  5  6  7  8
//  ---------------------------------
//  positions  1  0  0  2 10  4 40 92
//  ---------------------------------
//
// Author: Andreas Atle, atle.andreas@gmail.com
package queen

// Empty used to define a set
type Empty struct{}

// BoardSet contains a set of integers
type BoardSet map[int]Empty

// Board contains the moves on a chessboard of size N
type Board struct {
	N int
	B BoardSet
}

// NewBoard creates a new instance of a Board
func NewBoard(n int) *Board {
	b := &Board{N: n}
	b.B = BoardSet{}

	for pos := 0; pos < b.N*b.N; pos++ {
		b.B[pos] = Empty{}
	}
	return b
}

// CountAllPositions counts all possible configurations on a board
func (b *Board) CountAllPositions() int {
	count := 0
	for col := 0; col < b.N; col++ {
		b2 := b.copyTo()
		count += b.countPositionsStartingAt(0, col, 0)
		b.copyFrom(b2)
	}
	return count
}

// Pos returns the position in the slice for a given (row,col)-pair
func (b *Board) Pos(row, col int) int {
	return row*b.N + col
}

// prune eliminates the moves covered by a queen at (row,col)
func (b *Board) prune(row, col int) {

	// prune current row
	for c := 0; c < b.N; c++ {
		delete(b.B, b.Pos(row, c))
	}

	// prune current column
	for r := 0; r < b.N; r++ {
		delete(b.B, b.Pos(r, col))
	}

	for d := 1; d < b.N; d++ {
		if row-d >= 0 && col-d >= 0 {
			delete(b.B, b.Pos(row-d, col-d))
		}
		if row-d >= 0 && col+d < b.N {
			delete(b.B, b.Pos(row-d, col+d))
		}
		if row+d < b.N && col-d >= 0 {
			delete(b.B, b.Pos(row+d, col-d))
		}
		if row+d < b.N && col+d < b.N {
			delete(b.B, b.Pos(row+d, col+d))
		}
	}
}

// countPositionsStartingAt count the possible configs when (row,col) is included
func (b *Board) countPositionsStartingAt(row, col, nQueens int) int {
	// Assume that (row,col) is still valid
	if _, ok := b.B[b.Pos(row, col)]; !ok {
		return 0
	}

	// Put queen at (row,col) and prune moves
	nQueens++
	b.prune(row, col)

	// Check if done
	if nQueens == b.N {
		return 1
	}

	count := 0
	// Recursion over remaining moves
	b2 := b.copyTo()
	for c := 0; c < b.N; c++ {
		b.copyFrom(b2)
		count += b.countPositionsStartingAt(row+1, c, nQueens)
	}
	return count

}

// copyTo create a new instance of the Boardset in the Board
func (b *Board) copyTo() BoardSet {
	b2 := BoardSet{}
	for k := range b.B {
		b2[k] = Empty{}
	}
	return b2
}

// copyFrom copies back the copy of the BoardSet to the Board
func (b *Board) copyFrom(b2 BoardSet) {

	b.B = BoardSet{}
	for k := range b2 {
		b.B[k] = Empty{}
	}
}
