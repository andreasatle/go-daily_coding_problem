// Package islands contains Daily Coding Problem #84
// This problem was asked by Amazon.

// Given a Matrix of 1s and 0s, return the number of "islands"
// in the Matrix. A 1 represents land and 0 represents water,
// so an island is a group of 1s that are neighboring whose
// perimeter is surrounded by water.
// For example, this Matrix has 4 islands.

// 1 0 0 0 0
// 0 0 1 1 0
// 0 1 1 0 0
// 0 0 0 0 0
// 1 1 0 0 1
// 1 1 0 0 1

package islands

import "fmt"

// Entry is an int
type Entry int

// Matrix containing 0 and 1
type Matrix struct {
	Rows   int
	Cols   int
	Matrix []Entry
}

// Number of entries in Matrix
func (m *Matrix) size() int {
	return m.Rows * m.Cols
}

// Value at position (row, col) in Matrix
func (m *Matrix) getValue(row, col int) Entry {
	return m.Matrix[row*m.Cols+col]
}

// Set value at position (row, col) in Matrix
func (m *Matrix) setValue(row, col int, value Entry) {
	m.Matrix[row*m.Cols+col] = value
}

// Row for index in Matrix
func (m *Matrix) getRow(index int) int {
	return index / m.Cols
}

// Column for index in Matrix
func (m *Matrix) getCol(index int) int {
	return index % m.Cols
}

// Print the Matrix entries
func (m *Matrix) Print() {
	for row := 0; row < m.Rows; row++ {
		for col := 0; col < m.Cols; col++ {
			fmt.Print(m.getValue(row, col))
		}
		fmt.Println()
	}
}

// Clone a Matrix
func (m *Matrix) clone() *Matrix {
	new := &Matrix{m.Rows, m.Cols, []Entry{}}
	new.Matrix = make([]Entry, m.size())
	copy(new.Matrix, m.Matrix)
	return new
}

// Prune the current island recursively
func (m *Matrix) prune(row, col int) {
	if row >= 0 && row < m.Rows && col >= 0 && col < m.Cols {
		m.setValue(row, col, 0)
	}

	// Keep track of boundary of matrix
	avoidTopRow := row > 0
	avoidBottomRow := row < m.Rows-1
	avoidLeftCol := col > 0
	avoidRightCol := col < m.Cols-1

	if avoidTopRow {
		if avoidLeftCol && m.getValue(row-1, col-1) == 1 {
			m.prune(row-1, col-1)
		}
		if m.getValue(row-1, col) == 1 {
			m.prune(row-1, col)
		}
		if avoidRightCol && m.getValue(row-1, col+1) == 1 {
			m.prune(row-1, col+1)
		}
	}

	if avoidLeftCol && m.getValue(row, col-1) == 1 {
		m.prune(row, col-1)
	}
	if avoidRightCol && m.getValue(row, col+1) == 1 {
		m.prune(row, col+1)
	}

	if avoidBottomRow {
		if avoidLeftCol && m.getValue(row+1, col-1) == 1 {
			m.prune(row+1, col-1)
		}
		if m.getValue(row+1, col) == 1 {
			m.prune(row+1, col)
		}
		if avoidRightCol && m.getValue(row+1, col+1) == 1 {
			m.prune(row+1, col+1)
		}
	}
}

// CountIslands returns the number of isolated islands consisting of 1,
// that is separated by 0's.
func (m *Matrix) CountIslands() int {
	cnt := 0
	mat := m.clone()

	for row := 0; row < mat.Rows; row++ {
		for col := 0; col < mat.Cols; col++ {
			if mat.getValue(row, col) == 1 {
				cnt++
				mat.prune(row, col)
			}
		}
	}
	return cnt
}
