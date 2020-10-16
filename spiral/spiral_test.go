package spiral_test

import (
	"errors"
	"testing"

	"github.com/andreasatle/go-daily_coding_problem/spiral"
	"github.com/stretchr/testify/assert"
)

func TestSpiralCase1(t *testing.T) {

	mat := [][]int{[]int{1, 2, 3, 4, 5}, []int{6, 7, 8, 9, 10}, []int{11, 12, 13, 14, 15}, []int{16, 17, 18, 19, 20}}
	ref := []int{1, 2, 3, 4, 5, 10, 15, 20, 19, 18, 17, 16, 11, 6, 7, 8, 9, 14, 13, 12}
	slice, err := spiral.ToSlice(mat)

	assert.Equal(t, nil, err)
	for i := 0; i < len(ref); i++ {
		assert.Equal(t, ref[i], slice[i])
	}
}

func TestSpiralNonRect(t *testing.T) {

	mat := [][]int{[]int{1, 2, 3, 4, 5}, []int{6, 7, 8, 9, 10}, []int{11, 12, 13, 14, 15}, []int{16, 17, 18, 19}}
	_, err := spiral.ToSlice(mat)

	assert.Equal(t, errors.New("non-rectangular matrix"), err)
}

func TestSpiralEmpty(t *testing.T) {

	mat := [][]int{}
	slice, _ := spiral.ToSlice(mat)

	assert.Equal(t, []int{}, slice)
}
