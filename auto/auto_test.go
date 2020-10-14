package auto_test

import (
	"testing"

	"github.com/andreasatle/go-daily_coding_problem/auto"
	"github.com/stretchr/testify/assert"
)

func TestAuto(t *testing.T) {
	a := auto.NewAutoComplete([]string{"dog", "deer", "deal", "eel"})
	var c []string

	// Almost testcase from Daily Coding Problem
	c = a.Complete("de")
	assert.Equal(t, 2, len(c))
	assert.Equal(t, "deer", c[0])
	assert.Equal(t, "deal", c[1])

	// Should contain all words starting with d
	c = a.Complete("d")
	assert.Equal(t, 3, len(c))
	assert.Equal(t, "dog", c[0])
	assert.Equal(t, "deer", c[1])
	assert.Equal(t, "deal", c[2])

	// Should contain all words starting with e
	c = a.Complete("e")
	assert.Equal(t, 1, len(c))
	assert.Equal(t, "eel", c[0])

	// Should contain all words
	c = a.Complete("")
	assert.Equal(t, 4, len(c))
	assert.Equal(t, "dog", c[0])
	assert.Equal(t, "deer", c[1])
	assert.Equal(t, "deal", c[2])
	assert.Equal(t, "eel", c[3])

	// Should be empty
	c = a.Complete("f")
	assert.Equal(t, 0, len(c))
}
