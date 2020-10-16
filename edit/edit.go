// Package edit contains Daily Coding Problem #31
//
// This problem was asked by Google.
//
// The edit distance between two strings refers to the minimum number of
// character insertions, deletions, and substitutions required to change
// one string to the other. For example, the edit distance between “kitten”
// and “sitting” is three: substitute the “k” for “s”, substitute the “e”
// for “i”, and append a “g”.
//
// Given two strings, compute the edit distance between them.
//
// Note: We want to compute the Levenshtein-distance, and use a dynamic
// programming approach, i.e. Wagner–Fischer algorithm.
//
// Author: Andreas Atle, atle.andreas@gmail.com
package edit

type work struct {
	rows int
	cols int
	from []rune
	to   []rune
	mat  [][]int
}

func Distance(from, to []rune) int {
	w := initWork(from, to)
	return w.mat[w.rows-1][w.cols-1]
}

func initWork(from, to []rune) *work {
	w := &work{}
	w.rows = len(from) + 1
	w.cols = len(to) + 1
	w.from = from
	w.to = to
	w.initMat()
	return w
}

func (w *work) initMat() {
	// Allocate matrix
	w.mat = make([][]int, w.rows)
	for i := range w.mat {
		w.mat[i] = make([]int, w.cols)
	}

	// Initialize first column
	for row := 0; row < w.rows; row++ {
		w.mat[row][0] = row
	}

	// Initialize first row
	for col := 0; col < w.cols; col++ {
		w.mat[0][col] = col
	}

	needSubst := func(iFrom, iTo int) int {
		if w.from[iFrom] != w.to[iTo] {
			return 1
		}
		return 0
	}

	for row := 1; row < w.rows; row++ {
		for col := 1; col < w.cols; col++ {
			deleteCost := w.mat[row-1][col] + 1
			insertCost := w.mat[row][col-1] + 1
			substCost := w.mat[row-1][col-1] + needSubst(row-1, col-1)
			w.mat[row][col] = min(min(deleteCost, insertCost), substCost)
		}
	}
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
