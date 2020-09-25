package queen_test

import (
	"testing"

	"github.com/andreasatle/go-daily_coding_problem/queen"
	"github.com/stretchr/testify/assert"
)

func TestQueen(t *testing.T) {
	assert.Equal(t, 1, queen.NewBoard(1).CountAllPositions())
	assert.Equal(t, 0, queen.NewBoard(2).CountAllPositions())
	assert.Equal(t, 0, queen.NewBoard(3).CountAllPositions())
	assert.Equal(t, 2, queen.NewBoard(4).CountAllPositions())
	assert.Equal(t, 10, queen.NewBoard(5).CountAllPositions())
	assert.Equal(t, 4, queen.NewBoard(6).CountAllPositions())
	assert.Equal(t, 40, queen.NewBoard(7).CountAllPositions())
	assert.Equal(t, 92, queen.NewBoard(8).CountAllPositions())
}
