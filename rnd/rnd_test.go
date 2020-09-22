package rnd_test

import (
	"testing"

	"github.com/andreasatle/go-daily_coding_problem/rnd"

	"github.com/stretchr/testify/assert"
)

func TestRnd(t *testing.T) {
	// Check that error occurs if no valid values
	_, err := rnd.Create(5, []int{0, 1, 2, 3, 4})
	assert.NotEqual(t, nil, err)

	rand, _ := rnd.Create(5, []int{2, 4})

	// Randomize 1000 times and assume that 0, 1 and 3
	// are used atleast 250 times each.
	m := map[int]int{}
	for i := 0; i < 1000; i++ {
		r := rand()
		assert.NotEqual(t, true, r == 2 || r == 4)
		assert.Equal(t, true, r == 0 || r == 1 || r == 3)
		m[r]++
	}
	assert.Equal(t, true, m[0] >= 250)
	assert.Equal(t, true, m[1] >= 250)
	assert.Equal(t, true, m[3] >= 250)
}
