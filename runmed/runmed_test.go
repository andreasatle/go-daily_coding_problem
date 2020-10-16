package runmed_test

import (
	"testing"

	"github.com/andreasatle/go-daily_coding_problem/runmed"
	"github.com/stretchr/testify/assert"
)

func TestRunmedCase1(t *testing.T) {
	// Insert linearly increasing values
	seq := runmed.NewSequence()
	assert.Equal(t, 1.0, seq.Insert(1.0))
	assert.Equal(t, 1.5, seq.Insert(2.0))
	assert.Equal(t, 2.0, seq.Insert(3.0))
	assert.Equal(t, 2.5, seq.Insert(4.0))
	assert.Equal(t, 3.0, seq.Insert(5.0))
	assert.Equal(t, 3.5, seq.Insert(6.0))
	assert.Equal(t, 4.0, seq.Insert(7.0))
	assert.Equal(t, 4.5, seq.Insert(8.0))
}

func TestRunmedCase2(t *testing.T) {
	// Case from formulation
	seq := runmed.NewSequence()
	assert.Equal(t, 2.0, seq.Insert(2.0))
	assert.Equal(t, 1.5, seq.Insert(1.0))
	assert.Equal(t, 2.0, seq.Insert(5.0))
	assert.Equal(t, 3.5, seq.Insert(7.0))
	assert.Equal(t, 2.0, seq.Insert(2.0))
	assert.Equal(t, 2.0, seq.Insert(0.0))
	assert.Equal(t, 2.0, seq.Insert(5.0))
}
