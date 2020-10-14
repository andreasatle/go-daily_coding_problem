package sumprimes_test

import (
	"errors"
	"testing"

	"github.com/andreasatle/go-daily_coding_problem/sumprimes"
	"github.com/stretchr/testify/assert"
)

func help(n int) (int, int, int, error) {
	p1, p2, err := sumprimes.PrimeSum(n)
	return n, p1, p2, err
}
func TestSumPrimes(t *testing.T) {
	var p1, p2, n int
	var err error

	n, p1, p2, err = help(4)
	assert.Equal(t, 2, p1)
	assert.Equal(t, 2, p2)
	assert.Equal(t, n, p1+p2)

	n, p1, p2, err = help(6)
	assert.Equal(t, 3, p1)
	assert.Equal(t, 3, p2)
	assert.Equal(t, n, p1+p2)

	n, p1, p2, err = help(8)
	assert.Equal(t, 3, p1)
	assert.Equal(t, 5, p2)
	assert.Equal(t, n, p1+p2)

	n, p1, p2, err = help(10)
	assert.Equal(t, 3, p1)
	assert.Equal(t, 7, p2)
	assert.Equal(t, n, p1+p2)

	n, p1, p2, err = help(12)
	assert.Equal(t, 5, p1)
	assert.Equal(t, 7, p2)
	assert.Equal(t, n, p1+p2)

	n, p1, p2, err = help(14)
	assert.Equal(t, 3, p1)
	assert.Equal(t, 11, p2)
	assert.Equal(t, n, p1+p2)

	n, p1, p2, err = help(16)
	assert.Equal(t, 3, p1)
	assert.Equal(t, 13, p2)
	assert.Equal(t, n, p1+p2)

	n, p1, p2, err = help(11)
	assert.Equal(t, errors.New("odd input not allowed"), err)

	n, p1, p2, err = help(0)
	assert.Equal(t, errors.New("no primes found"), err)
}
