package stair_test

import (
	"fmt"
	"testing"

	"github.com/andreasatle/go-daily_coding_problem/stair"
	"github.com/stretchr/testify/assert"
)

func TestStair(t *testing.T) {
	assert.Equal(t, 5, len(stair.UniqueWays(4, []int{1, 2}, []int{}, [][]int{})))
}

func Example_stair() {
	lst := stair.UniqueWays(4, []int{1, 2}, []int{}, [][]int{})
	fmt.Println(lst)
	// output:
	// [[1 1 1 1] [1 1 2] [1 2 1] [2 1 1] [2 2]]
}
