package letters_test

import (
	"testing"

	"github.com/andreasatle/go-daily_coding_problem/letters"
	"github.com/stretchr/testify/assert"
)

func TestLetters(t *testing.T) {
	board, _ := letters.NewBoard([]string{"ABCE", "SFCS", "ADEE"})
	assert.Equal(t, true, board.PathExists("ABCCED"))
	assert.Equal(t, true, board.PathExists("SEE"))
	assert.Equal(t, false, board.PathExists("ABCB"))
	assert.Equal(t, true, board.PathExists("BCCF"))
	assert.Equal(t, false, board.PathExists("BCCFB"))
}
