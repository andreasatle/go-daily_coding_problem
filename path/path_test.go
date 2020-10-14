package path_test

import (
	"errors"
	"testing"

	"github.com/andreasatle/go-daily_coding_problem/path"
	"github.com/stretchr/testify/assert"
)

func TestPath(t *testing.T) {
	// Funky way of defining the graph with a cascaded call
	length, err := path.NewVertex("ABACA").AddEdge(0, 1).AddEdge(0, 2).AddEdge(2, 3).AddEdge(3, 4).LargestValue()
	assert.Equal(t, 3, length)

	// Check for cyclic graph
	length, err = path.NewVertex("ABACA").AddEdge(0, 1).AddEdge(1, 0).LargestValue()
	assert.Equal(t, errors.New("cyclic graph"), err)

	// Funky way of defining the graph with a cascaded call
	length, err = path.NewVertex("ABACA").AddEdge(0, 1).AddEdge(0, 2).AddEdge(2, 3).AddEdge(3, 4).AddEdge(4, 3).LargestValue()
	assert.Equal(t, errors.New("cyclic graph"), err)
}
