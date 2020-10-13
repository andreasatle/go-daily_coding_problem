package stringshift_test

import (
	"testing"

	"github.com/andreasatle/go-daily_coding_problem/stringshift"
	"github.com/stretchr/testify/assert"
)

func TestStringshift(t *testing.T) {
	// Rotation exists
	assert.True(t, stringshift.ShiftedEqual("abcde", "cdeab"))

	// No rotation exists
	assert.False(t, stringshift.ShiftedEqual("abc", "acb"))

	// Check empty strings
	assert.True(t, stringshift.ShiftedEqual("", ""))

	// Check equality of strings of size 1
	assert.True(t, stringshift.ShiftedEqual("Q", "Q"))

	// Check inequality of strings of size 1
	assert.False(t, stringshift.ShiftedEqual("Q", "q"))

	// Check strings of different sizes
	assert.False(t, stringshift.ShiftedEqual("", "r"))
}
