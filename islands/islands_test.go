package islands_test

import (
	"testing"

	"github.com/andreasatle/go-daily_coding_problem/islands"
	"github.com/stretchr/testify/assert"
)

func TestIslands(t *testing.T) {
	var mat *islands.Matrix

	// Test with the case given by "Daily Coding Problem"
	mat = &islands.Matrix{6, 5, []islands.Entry{
		1, 0, 0, 0, 0,
		0, 0, 1, 1, 0,
		0, 1, 1, 0, 0,
		0, 0, 0, 0, 0,
		1, 1, 0, 0, 1,
		1, 1, 0, 0, 1,
	}}
	assert.Equal(t, 4, mat.CountIslands(), "Wrong number of islands.")

	// Add zeros around the previous test case
	mat = &islands.Matrix{8, 7, []islands.Entry{
		0, 0, 0, 0, 0, 0, 0,
		0, 1, 0, 0, 0, 0, 0,
		0, 0, 0, 1, 1, 0, 0,
		0, 0, 1, 1, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0,
		0, 1, 1, 0, 0, 1, 0,
		0, 1, 1, 0, 0, 1, 0,
		0, 0, 0, 0, 0, 0, 0,
	}}
	assert.Equal(t, 4, mat.CountIslands(), "Wrong number of islands.")
}
