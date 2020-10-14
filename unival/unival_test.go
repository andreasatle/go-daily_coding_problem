package unival_test

import (
	"testing"

	"github.com/andreasatle/go-daily_coding_problem/unival"
	"github.com/stretchr/testify/assert"
)

func TestUnival(t *testing.T) {
	tree := &unival.Node{0, &unival.Node{1, nil, nil}, &unival.Node{0, &unival.Node{1, &unival.Node{1, nil, nil}, &unival.Node{1, nil, nil}}, &unival.Node{0, nil, nil}}}
	assert.Equal(t, 5, tree.Count())
}

func TestUnivalEmpty(t *testing.T) {
	var tree *unival.Node
	assert.Equal(t, 0, tree.Count())
}
