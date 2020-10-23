package corner_test

import (
	"testing"

	"github.com/andreasatle/go-daily_coding_problem/corner"
	"github.com/stretchr/testify/assert"
)

func TestCorner2x2(t *testing.T) {
	assert.Equal(t, 2, corner.CountPathsV1(2, 2), "2x2 grid has 2 paths, recursion failed")
	assert.Equal(t, 2, corner.CountPathsV2(2, 2), "2x2 grid has 2 paths, dynamic programming failed")
}

func TestCorner0x2(t *testing.T) {
	assert.Equal(t, 0, corner.CountPathsV1(0, 2), "0x2 grid has 0 paths, recursion failed")
	assert.Equal(t, 0, corner.CountPathsV2(0, 2), "0x2 grid has 0 paths, dynamic programming failed")
}

func TestCorner3x3(t *testing.T) {
	assert.Equal(t, 6, corner.CountPathsV1(3, 3), "3x3 grid has 6 paths, recursion failed")
	assert.Equal(t, 6, corner.CountPathsV2(3, 3), "3x3 grid has 6 paths, dynamic programming failed")
}

func TestCorner5x5(t *testing.T) {
	assert.Equal(t, 70, corner.CountPathsV1(5, 5), "5x5 grid has 70 paths, recursion failed")
	assert.Equal(t, 70, corner.CountPathsV2(5, 5), "5x5 grid has 70 paths, dynamic programming failed")
}
