package sumprimes_test

import (
	"testing"

	"github.com/andreasatle/go-daily_coding_problem/sumprimes"
	"github.com/stretchr/testify/assert"
)

func help(n int) (int, int, int) {
	p1, p2 := sumprimes.PrimeSum(n)
	return n, p1, p2
}
func TestSumPrimes(t *testing.T) {
	var p1, p2, n int

	n, p1, p2 = help(4)
	assert.Equal(t, 2, p1)
	assert.Equal(t, 2, p2)
	assert.Equal(t, n, p1+p2)

	n, p1, p2 = help(6)
	assert.Equal(t, 3, p1)
	assert.Equal(t, 3, p2)
	assert.Equal(t, n, p1+p2)

	n, p1, p2 = help(8)
	assert.Equal(t, 3, p1)
	assert.Equal(t, 5, p2)
	assert.Equal(t, n, p1+p2)

	n, p1, p2 = help(10)
	assert.Equal(t, 3, p1)
	assert.Equal(t, 7, p2)
	assert.Equal(t, n, p1+p2)

	n, p1, p2 = help(12)
	assert.Equal(t, 5, p1)
	assert.Equal(t, 7, p2)
	assert.Equal(t, n, p1+p2)

	n, p1, p2 = help(14)
	assert.Equal(t, 3, p1)
	assert.Equal(t, 11, p2)
	assert.Equal(t, n, p1+p2)

	n, p1, p2 = help(16)
	assert.Equal(t, 3, p1)
	assert.Equal(t, 13, p2)
	assert.Equal(t, n, p1+p2)
}
