package negsort_test

import (
	"testing"

	"github.com/andreasatle/go-daily_coding_problem/negsort"
	"github.com/stretchr/testify/assert"
)

func TestNegsortPosAndNeg(t *testing.T) {
	vals := []int{-7, -6, -4, -1, 3, 4, 5}
	negsort.SortSquares(vals)
	assert.Equal(t, len(vals), 7)
	assert.Equal(t, 1, vals[0])
	assert.Equal(t, 9, vals[1])
	assert.Equal(t, 16, vals[2])
	assert.Equal(t, 16, vals[3])
	assert.Equal(t, 25, vals[4])
	assert.Equal(t, 36, vals[5])
	assert.Equal(t, 49, vals[6])
}

func TestNegsortEmpty(t *testing.T) {
	vals := []int{}
	negsort.SortSquares(vals)
	assert.Equal(t, len(vals), 0)
}

func TestNegsortPositive(t *testing.T) {
	vals := []int{1, 3, 6, 8, 9}
	negsort.SortSquares(vals)
	assert.Equal(t, len(vals), 5)
	assert.Equal(t, 1, vals[0])
	assert.Equal(t, 9, vals[1])
	assert.Equal(t, 36, vals[2])
	assert.Equal(t, 64, vals[3])
	assert.Equal(t, 81, vals[4])
}

func TestNegsortNegative(t *testing.T) {
	vals := []int{-9, -8, -6, -3, -1}
	negsort.SortSquares(vals)
	assert.Equal(t, len(vals), 5)
	assert.Equal(t, 1, vals[0])
	assert.Equal(t, 9, vals[1])
	assert.Equal(t, 36, vals[2])
	assert.Equal(t, 64, vals[3])
	assert.Equal(t, 81, vals[4])
}
