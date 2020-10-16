package edit_test

import (
	"testing"

	"github.com/andreasatle/go-daily_coding_problem/edit"
	"github.com/stretchr/testify/assert"
)

func TestCase1(t *testing.T) {
	assert.Equal(t, 3, edit.Distance([]rune("Sunday"), []rune("Saturday")))
	assert.Equal(t, 3, edit.Distance([]rune("Saturday"), []rune("Sunday")))
}

func TestCase2(t *testing.T) {
	assert.Equal(t, 3, edit.Distance([]rune("kitten"), []rune("sitting")))
	assert.Equal(t, 3, edit.Distance([]rune("sitting"), []rune("kitten")))
}

func TestCase3(t *testing.T) {
	assert.Equal(t, 5, edit.Distance([]rune(""), []rune("12345")))
	assert.Equal(t, 5, edit.Distance([]rune("12345"), []rune("")))
}

func TestCase4(t *testing.T) {
	assert.Equal(t, 2, edit.Distance([]rune("xabcd"), []rune("abcdx")))
	assert.Equal(t, 2, edit.Distance([]rune("abcdx"), []rune("xabcd")))
}
