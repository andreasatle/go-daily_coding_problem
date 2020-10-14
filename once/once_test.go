package once_test

import (
	"testing"

	"github.com/andreasatle/go-daily_coding_problem/once"
	"github.com/stretchr/testify/assert"
)

func TestOnce(t *testing.T) {
	var slice []int

	slice = []int{1, 1, 3, 5, 3, 1, 3}
	assert.Equal(t, 5, once.Find(slice))

	slice = []int{7, -8, 3, 4, 5, 7, -8, 3, 6, 4, 5, 7, -8, 3, 4, 5}
	assert.Equal(t, 6, once.Find(slice))

	slice = []int{4711}
	assert.Equal(t, 4711, once.Find(slice))

	slice = []int{}
	assert.Equal(t, 0, once.Find(slice))

	slice = []int{-7, 0, -7, -7}
	assert.Equal(t, 0, once.Find(slice))

	slice = []int{0, -3, 0, 0}
	assert.Equal(t, -3, once.Find(slice))
}
