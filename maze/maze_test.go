package maze_test

import (
	"testing"

	"github.com/andreasatle/go-daily_coding_problem/maze"
	"github.com/stretchr/testify/assert"
)

func TestMazeCase1(t *testing.T) {
	// Case from problem formulation
	board := [][]bool{
		[]bool{false, false, false, false},
		[]bool{true, true, false, true},
		[]bool{false, false, false, false},
		[]bool{false, false, false, false},
	}

	start := maze.Coord{Row: 3, Col: 0}
	end := maze.Coord{Row: 0, Col: 0}
	shortestPath, ok := maze.PathLength(board, start, end)
	assert.True(t, ok)
	assert.Equal(t, 7, shortestPath)
}

func TestMazeCase2(t *testing.T) {
	// Spiral case
	board := [][]bool{
		[]bool{false, false, false, false, false},
		[]bool{true, true, true, true, false},
		[]bool{false, false, false, true, false},
		[]bool{false, true, true, true, false},
		[]bool{false, false, false, false, false},
	}

	start := maze.Coord{Row: 0, Col: 0}
	end := maze.Coord{Row: 2, Col: 2}
	shortestPath, ok := maze.PathLength(board, start, end)
	assert.True(t, ok)
	assert.Equal(t, 16, shortestPath)
}

func TestMazeCase3(t *testing.T) {
	// Source at wall
	board := [][]bool{
		[]bool{false, false, false, false, false},
		[]bool{true, true, true, true, false},
		[]bool{false, false, false, true, false},
		[]bool{false, true, true, true, false},
		[]bool{false, false, false, false, false},
	}

	start := maze.Coord{Row: 1, Col: 0}
	end := maze.Coord{Row: 2, Col: 2}
	_, ok := maze.PathLength(board, start, end)
	assert.False(t, ok)
}

func TestMazeCase4(t *testing.T) {
	// Target at wall
	board := [][]bool{
		[]bool{false, false, false, false, false},
		[]bool{true, true, true, true, false},
		[]bool{false, false, false, true, false},
		[]bool{false, true, true, true, false},
		[]bool{false, false, false, false, false},
	}

	start := maze.Coord{Row: 0, Col: 0}
	end := maze.Coord{Row: 1, Col: 0}
	_, ok := maze.PathLength(board, start, end)
	assert.False(t, ok)
}
