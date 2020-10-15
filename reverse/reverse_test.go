package reverse_test

import (
	"testing"

	"github.com/andreasatle/go-daily_coding_problem/reverse"
	"github.com/stretchr/testify/assert"
)

func TestReverse1(t *testing.T) {
	words := []rune("here I golang  again")
	reverse.Reverse(words)
	assert.Equal(t, "again  golang I here", string(words))
}

func TestReverse2(t *testing.T) {
	words := []rune("")
	reverse.Reverse(words)
	assert.Equal(t, "", string(words))
}

func TestReverse3(t *testing.T) {
	words := []rune(" ")
	reverse.Reverse(words)
	assert.NotEqual(t, "  ", string(words))
}

func TestReverse4(t *testing.T) {
	words := []rune(" foo bar")
	reverse.Reverse(words)
	assert.Equal(t, "bar foo ", string(words))
}
