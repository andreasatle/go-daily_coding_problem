package nonadj_test

import (
	"testing"

	"github.com/andreasatle/go-daily_coding_problem/nonadj"
	"github.com/stretchr/testify/assert"
)

func TestNonadj(t *testing.T) {

	// Test case from daily coding problem
	assert.Equal(t, 13, nonadj.Sum([]int{2, 4, 6, 2, 5}))

	// Test case from daily coding problem
	assert.Equal(t, 10, nonadj.Sum([]int{5, 1, 1, 5}))

	// Check empty array
	assert.Equal(t, 0, nonadj.Sum([]int{}))

	// Check all entries negative
	assert.Equal(t, 0, nonadj.Sum([]int{-1, -2, -3}))

	// Some more examples
	assert.Equal(t, 13, nonadj.Sum([]int{0, 2, -3, 4, 1, -1, 7}))

	// Some more examples
	assert.Equal(t, 13, nonadj.Sum([]int{0, 2, -3, 1, 4, -1, 7}))
}
