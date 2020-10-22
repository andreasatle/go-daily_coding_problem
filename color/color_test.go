package color_test

import (
	"testing"

	"github.com/andreasatle/go-daily_coding_problem/color"
	"github.com/stretchr/testify/assert"
)

func TestColorCase1(t *testing.T) {
	// Check where last color-index is 0
	cost := [][]float64{
		{1.0, 2.0, 3.0},
		{1.0, 3.0, 4.0},
		{2.0, 2.0, 4.0},
		{2.0, 2.0, 4.0},
	}
	totalCost, colors := color.Color(cost)
	assert.Equal(t, totalCost, 7.0, "Total cost is wrong")
	assert.Equal(t, 1, colors[0], "First color is wrong")
	assert.Equal(t, 0, colors[1], "Second color is wrong")
	assert.Equal(t, 1, colors[2], "Third color is wrong")
	assert.Equal(t, 0, colors[3], "Fourth color is wrong")
}

func TestColorCase2(t *testing.T) {
	// Check where last color-index is 2, for full coverage
	cost := [][]float64{
		{3.0, 2.0, 1.0},
		{4.0, 3.0, 1.0},
		{4.0, 2.0, 2.0},
		{4.0, 2.0, 2.0},
	}
	totalCost, colors := color.Color(cost)
	assert.Equal(t, totalCost, 7.0, "Total cost is wrong")
	assert.Equal(t, 1, colors[0], "First color is wrong")
	assert.Equal(t, 2, colors[1], "Second color is wrong")
	assert.Equal(t, 1, colors[2], "Third color is wrong")
	assert.Equal(t, 2, colors[3], "Fourth color is wrong")
}

func TestColorCase3(t *testing.T) {
	// Check trivial empty problem
	cost := [][]float64{}
	totalCost, colors := color.Color(cost)
	assert.Equal(t, totalCost, 0.0, "Total cost is wrong")
	assert.Equal(t, 0, len(colors), "Color-indices are non-empty")
}
