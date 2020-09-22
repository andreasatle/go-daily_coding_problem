package subseq_test

import (
	"strconv"
	"testing"

	"github.com/andreasatle/go-daily_coding_problem/subseq"
	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	array := []int{0, 8, 4, 12, 2, 10, 6, 14, 1, 9, 5, 13, 3, 11, 7, 15}
	out := subseq.Subsequence(array, []int{})
	ref := []int{0, 2, 6, 9, 11}
	assert.Equal(t, 5, len(out), "Length of output not correct.")
	for i := 0; i < 5; i++ {
		assert.Equal(t, ref[i], out[i], "Index "+strconv.Itoa(i)+" is wrong")
	}
}
