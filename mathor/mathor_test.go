package mathor_test

import (
	"testing"

	"github.com/andreasatle/go-daily_coding_problem/mathor"
	"github.com/stretchr/testify/assert"
)

func TestMathor(t *testing.T) {
	var x, y int
	x = 4711
	y = -3711
	x2, _ := mathor.Ternary(x, y, 1)
	y2, _ := mathor.Ternary(x, y, 0)
	_, err := mathor.Ternary(x, y, 2)
	assert.Equal(t, x, x2)
	assert.Equal(t, y, y2)
	assert.NotEqual(t, nil, err)
}
