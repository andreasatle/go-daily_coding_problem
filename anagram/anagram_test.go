package anagram_test

import (
	"testing"

	"github.com/andreasatle/go-daily_coding_problem/anagram"
	"github.com/stretchr/testify/assert"
)

func TestAnagram(t *testing.T) {
	var indices []int

	indices = anagram.GetIndices("abxaba", "ab")
	assert.Equal(t, 3, len(indices))
	assert.Equal(t, 0, indices[0])
	assert.Equal(t, 3, indices[1])
	assert.Equal(t, 4, indices[2])
}
