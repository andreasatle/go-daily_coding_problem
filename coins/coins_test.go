package coins_test

import (
	"testing"

	"github.com/andreasatle/go-daily_coding_problem/coins"
	"github.com/stretchr/testify/assert"
)

func TestCoinsFromFormulation(t *testing.T) {
	assert.Equal(t, 12, coins.Max([][]int{
		[]int{0, 3, 1, 1},
		[]int{2, 0, 0, 4},
		[]int{1, 5, 3, 1},
	}), "problem from formulation should have answer 12")
}

func TestCoinsEmptyColumns(t *testing.T) {
	assert.Equal(t, 0, coins.Max([][]int{[]int{}, []int{}, []int{}}), "solution should be 0 for empty columns")
}

func TestCoinsEmptyMatrix(t *testing.T) {
	assert.Equal(t, 0, coins.Max([][]int{}), "solution should be 0 for empty matrix")
}
