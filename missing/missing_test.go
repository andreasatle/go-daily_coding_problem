package missing_test

import (
	"testing"

	"github.com/andreasatle/go-daily_coding_problem/missing"

	"github.com/stretchr/testify/assert"
)

func TestSumK(t *testing.T) {
	assert.Equal(t, 2, missing.FirstMissingPosInt([]int{3, 4, -1, 1}))
	assert.Equal(t, 1, missing.FirstMissingPosInt([]int{-1, -7, 2, -2}), "Value 1 is missing.")
	assert.Equal(t, 1, missing.FirstMissingPosInt([]int{-1, -7, -2, -2}), "No positive values.")
	assert.Equal(t, 4, missing.FirstMissingPosInt([]int{-1, -2, 3, 2, 1}), "First value not included is missing.")
}
