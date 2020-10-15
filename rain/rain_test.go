package rain_test

import (
	"testing"

	"github.com/andreasatle/go-daily_coding_problem/rain"
	"github.com/stretchr/testify/assert"
)

func TestRainCase1(t *testing.T) {
	// Case 1 from formulation
	assert.Equal(t, 1, rain.Fill([]int{2, 1, 2}))
}

func TestRainCase2(t *testing.T) {
	// Case 2 from formulation
	assert.Equal(t, 8, rain.Fill([]int{3, 0, 1, 3, 0, 5}))
}

func TestRainCase3(t *testing.T) {
	// Added some complexity to case from formulation
	assert.Equal(t, 9, rain.Fill([]int{1, 2, 0, 3, 0, 1, 4, 0, 2, 1}))
}

func TestRainCase4(t *testing.T) {
	// Case with no "outer walls"
	assert.Equal(t, 0, rain.Fill([]int{7, 31}))
}

func TestRainCase5(t *testing.T) {
	// Rather complicated example
	assert.Equal(t, 10, rain.Fill([]int{1, 2, 0, 3, 0, 1, 4, 2, 1, 3, 2}))
}

func TestRainCase6(t *testing.T) {
	// Think about complexity for decreasing case
	assert.Equal(t, 0, rain.Fill([]int{5, 4, 3, 3, 2, 1}))
}

func TestRainCase7(t *testing.T) {
	// Think about complexity for increasing case
	assert.Equal(t, 0, rain.Fill([]int{1, 2, 3, 3, 4, 5}))
}
