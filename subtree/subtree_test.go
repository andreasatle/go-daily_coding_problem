package subtree_test

import (
	"testing"

	"github.com/andreasatle/go-daily_coding_problem/subtree"
	"github.com/stretchr/testify/assert"
)

func TestSubtreeCase1(t *testing.T) {
	// Create two identical trees (but with different pointers)
	s2 := &subtree.Node{2, nil, nil}
	s3 := &subtree.Node{3, nil, nil}
	s1 := &subtree.Node{1, s2, s3}

	t2 := &subtree.Node{2, nil, nil}
	t3 := &subtree.Node{3, nil, nil}
	t1 := &subtree.Node{1, t2, t3}
	t12 := &subtree.Node{1, t2, nil}
	t13 := &subtree.Node{1, nil, t3}
	t1flip := &subtree.Node{1, t3, t2}

	assert.True(t, s1.IsSubtree(t1))
	assert.True(t, s1.IsSubtree(t2))
	assert.True(t, s1.IsSubtree(t3))
	assert.False(t, s1.IsSubtree(t12))
	assert.False(t, s1.IsSubtree(t13))
	assert.False(t, s1.IsSubtree(t1flip))

	assert.False(t, s2.IsSubtree(t1))
	assert.True(t, s2.IsSubtree(t2))
	assert.False(t, s2.IsSubtree(t3))

	assert.False(t, s3.IsSubtree(t1))
	assert.False(t, s3.IsSubtree(t2))
	assert.True(t, s3.IsSubtree(t3))
}

func TestSubtreeCase2(t *testing.T) {
	// Create two identical trees (with same pointers)
	s2 := &subtree.Node{2, nil, nil}
	s3 := &subtree.Node{3, nil, nil}
	s1 := &subtree.Node{1, s2, s3}

	assert.True(t, s1.IsSubtree(s1))
	assert.True(t, s1.IsSubtree(s2))
	assert.True(t, s1.IsSubtree(s3))

	assert.False(t, s2.IsSubtree(s1))
	assert.True(t, s2.IsSubtree(s2))
	assert.False(t, s2.IsSubtree(s3))

	assert.False(t, s3.IsSubtree(s1))
	assert.False(t, s3.IsSubtree(s2))
	assert.True(t, s3.IsSubtree(s3))
}
