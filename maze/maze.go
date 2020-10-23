// Package maze contains Daily Coding Problem #23
//
// This problem was asked by Google.
//
// You are given an M by N matrix consisting of booleans that represents
// a board. Each True boolean represents a wall. Each False boolean
// represents a tile you can walk on.
//
// Given this matrix, a start coordinate, and an end coordinate, return
// the minimum number of steps required to reach the end coordinate from
// the start. If there is no possible path, then return null. You can
// move up, left, down, and right. You cannot move through walls. You
// cannot wrap around the edges of the board.
//
// For example, given the following board:
//
//  [[f, f, f, f],
//   [t, t, f, t],
//   [f, f, f, f],
//   [f, f, f, f]]
// and start = (3, 0) (bottom left) and end = (0, 0) (top left), the
// minimum number of steps required to reach the end is 7, since we would
// need to go through (1, 2) because there is a wall everywhere else on
// the second row.
//
// Note: The problem is solved using backtracking. It is related to
// dynamic programming.
//
// Dynamic programming is more like BFS: we find all possible suboptimal
// solutions represented the non-leaf nodes, and only grow the tree by one
// layer under those non-leaf nodes.
//
// Backtracking is more like DFS: we grow the tree as deep as possible and
// prune the tree at one node if the solutions under the node are not what
// we expect.
//
// Author: Andreas Atle, atle.andreas@gmail.com
package maze

import (
	"math"
)

// Coord contains the indices to one entry in the board
type Coord struct {
	Row int
	Col int
}

// PathLength returns the minimum number of steps to reach the target
func PathLength(board [][]bool, start, target Coord) (int, bool) {
	n := len(board)
	// Check if start is out of bounds
	if start.Row < 0 || start.Row >= n || start.Col < 0 || start.Col >= len(board[start.Row]) {
		return 0, false
	}

	// Check that not a wall
	if board[start.Row][start.Col] {
		return 0, false
	}

	// Are we done?!
	if start == target {
		return 0, true
	}

	// Backtracking algorithm

	// Set board-value at the current position
	board[start.Row][start.Col] = true

	nSteps := math.MaxInt64

	up, ok := PathLength(board, Coord{Row: start.Row - 1, Col: start.Col}, target)
	nSteps = minIfOkElseFirstArg(nSteps, up, ok)

	down, ok := PathLength(board, Coord{Row: start.Row + 1, Col: start.Col}, target)
	nSteps = minIfOkElseFirstArg(nSteps, down, ok)

	left, ok := PathLength(board, Coord{Row: start.Row, Col: start.Col - 1}, target)
	nSteps = minIfOkElseFirstArg(nSteps, left, ok)

	right, ok := PathLength(board, Coord{Row: start.Row, Col: start.Col + 1}, target)
	nSteps = minIfOkElseFirstArg(nSteps, right, ok)

	// Unset the board-value at the current position
	board[start.Row][start.Col] = false

	// No direction is valid
	if nSteps == math.MaxInt64 {
		return -1, false
	}

	// Return minimum number of steps to target
	return 1 + nSteps, true
}

// Special min of int, where the first argument is returned when ok is false.
// This cleans up the code where its used.
func minIfOkElseFirstArg(a, b int, ok bool) int {
	if !ok || a <= b {
		return a
	}
	return b
}
