package knight_test

import (
	"testing"

	"github.com/andreasatle/go-daily_coding_problem/knight"
	"github.com/stretchr/testify/assert"
)

func TestKnight(t *testing.T) {

	assert.Equal(t, 0, knight.NewBoard(0).CountAllTours(), "Wrong number of knight moves for board-size 0.")
	assert.Equal(t, 1, knight.NewBoard(1).CountAllTours(), "Wrong number of knight moves for board-size 1.")
	assert.Equal(t, 0, knight.NewBoard(2).CountAllTours(), "Wrong number of knight moves for board-size 2.")
	assert.Equal(t, 0, knight.NewBoard(3).CountAllTours(), "Wrong number of knight moves for board-size 3.")
	assert.Equal(t, 0, knight.NewBoard(4).CountAllTours(), "Wrong number of knight moves for board-size 4.")
	assert.Equal(t, 1728, knight.NewBoard(5).CountAllTours(), "Wrong number of knight moves for board-size 5.")
}
