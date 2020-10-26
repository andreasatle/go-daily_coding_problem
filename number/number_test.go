package number_test

import (
	"testing"

	"github.com/andreasatle/go-daily_coding_problem/number"
	"github.com/stretchr/testify/assert"
)

func TestNumber(t *testing.T) {
	// Run the examples from the problem formulation
	assert.True(t, number.Check("10"))
	assert.True(t, number.Check("-10"))
	assert.True(t, number.Check("10.1"))
	assert.True(t, number.Check("-10.1"))
	assert.True(t, number.Check("1e5"))
	assert.True(t, number.Check("-1e5"))
	assert.True(t, number.Check("1E5"))
	assert.True(t, number.Check("-1E5"))
	assert.False(t, number.Check("a"))
	assert.False(t, number.Check("x 1"))
	assert.False(t, number.Check("a -2"))
	assert.False(t, number.Check("-"))
}
