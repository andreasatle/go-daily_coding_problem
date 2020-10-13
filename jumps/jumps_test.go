package jumps_test

import (
	"testing"

	"github.com/andreasatle/go-daily_coding_problem/jumps"
	"github.com/stretchr/testify/assert"
)

func TestJumps(t *testing.T) {
	// Last entry is reachable
	assert.True(t, jumps.FirstToLast([]int{2, 0, 1, 0}))

	// Can't get passed index 2
	assert.False(t, jumps.FirstToLast([]int{1, 1, 0, 1}))

	// Infinite loop at index 1-2-1-2-...
	assert.False(t, jumps.FirstToLast([]int{1, 1, -1, 1}))

	// Index too small
	assert.False(t, jumps.FirstToLast([]int{1, -2, -1, 1}))

	// Index too large
	assert.False(t, jumps.FirstToLast([]int{1, 1, 2, 1}))

	// Jumpy but ok (negative offset)
	assert.True(t, jumps.FirstToLast([]int{2, 2, -1, 1}))

	// Check that we don't abort too early (+-1 errors)
	assert.True(t, jumps.FirstToLast([]int{1, 1, 1, 1}))
}
