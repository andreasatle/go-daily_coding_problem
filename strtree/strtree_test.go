package strtree_test

import (
	"testing"

	"github.com/andreasatle/go-daily_coding_problem/strtree"
	"github.com/stretchr/testify/assert"
)

func TestTree(t *testing.T) {
	tree := &strtree.Node{"root", nil, nil}
	_ = tree
	assert.Equal(t, true, true)
}
