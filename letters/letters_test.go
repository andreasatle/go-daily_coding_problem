package letters_test

import (
	"testing"

	"github.com/andreasatle/go-daily_coding_problem/letters"
	"github.com/stretchr/testify/assert"
)

func TestLetters(t *testing.T) {
	var board *letters.Board
	var err error
	board, _ = letters.NewBoard([]string{"ABCE", "SFCS", "ADEE"})
	assert.True(t, board.PathExists("ABCCED"))
	assert.True(t, board.PathExists("SEE"))
	assert.False(t, board.PathExists("ABCB"))
	assert.True(t, board.PathExists("BCCF"))
	assert.False(t, board.PathExists("BCCFB"))
	assert.True(t, board.PathExists(""))

	_, err = letters.NewBoard([]string{})
	assert.False(t, err == nil)

	_, err = letters.NewBoard([]string{"ABCE", "FCS", "ADEE"})
	assert.False(t, err == nil)
}
