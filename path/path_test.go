package path_test

import (
	"testing"

	"github.com/andreasatle/go-daily_coding_problem/path"
	"github.com/stretchr/testify/assert"
)

func TestPath(t *testing.T) {
	// Funky way of defining the graph with a cascaded call
	g := path.NewVertex("ABACA").AddEdge(0, 1).AddEdge(0, 2).AddEdge(2, 3).AddEdge(3, 4)

	length, err := g.LargestValue()
	assert.Equal(t, 3, length)

	g = path.NewVertex("ABACA").AddEdge(0, 1).AddEdge(1, 0)
	length, err = g.LargestValue()
	assert.NotEqual(t, nil, err)

	g = path.NewVertex("ABACA").AddEdge(0, 0)
	length, err = g.LargestValue()
	assert.NotEqual(t, nil, err)
}
