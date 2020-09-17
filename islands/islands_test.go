package islands_test

import (
	"testing"

	"github.com/andreasatle/Assorted/Daily/islands"
	"github.com/stretchr/testify/assert"
)

func TestIslands(t *testing.T) {
	// Test with the case given by "Daily Coding Problem"
	mat := &islands.Matrix{6, 5, []islands.Entry{1, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0, 1}}
	assert.Equal(t, 4, mat.CountIslands(), "Wrong number of islands.")
}
