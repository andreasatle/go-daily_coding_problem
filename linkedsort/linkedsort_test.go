package linkedsort_test

import (
	"fmt"
	"testing"

	"github.com/andreasatle/go-daily_coding_problem/linkedsort"
	"github.com/stretchr/testify/assert"
)

func TestLinkedSortCase1(t *testing.T) {
	l1 := &linkedsort.LinkedList{}
	l1.Insert(7)
	l1.Insert(3)
	l1.Insert(10)
	l2 := &linkedsort.LinkedList{&linkedsort.Node{2, &linkedsort.Node{5, &linkedsort.Node{12, nil}}}}
	l3 := &linkedsort.LinkedList{&linkedsort.Node{4, &linkedsort.Node{5, &linkedsort.Node{8, nil}}}}
	l4 := &linkedsort.LinkedList{}
	l5 := &linkedsort.LinkedList{&linkedsort.Node{4, &linkedsort.Node{5, &linkedsort.Node{1, nil}}}}
	var l6 *linkedsort.LinkedList
	assert.True(t, l1.IsSorted())
	assert.True(t, l2.IsSorted())
	assert.True(t, l3.IsSorted())
	assert.True(t, l4.IsSorted())
	assert.False(t, l5.IsSorted())
	assert.False(t, l6.IsSorted())

	ll := linkedsort.Merge(l1, l2, l3, l4)
	llSlice := ll.Slice()

	assert.Equal(t, 2, llSlice[0])
	assert.Equal(t, 3, llSlice[1])
	assert.Equal(t, 4, llSlice[2])
	assert.Equal(t, 5, llSlice[3])
	assert.Equal(t, 5, llSlice[4])
	assert.Equal(t, 7, llSlice[5])
	assert.Equal(t, 8, llSlice[6])
	assert.Equal(t, 10, llSlice[7])
	assert.Equal(t, 12, llSlice[8])
}

func TestLinkedSortCase2(t *testing.T) {
	l1 := &linkedsort.LinkedList{}

	ll := linkedsort.Merge(l1)
	llSlice := ll.Slice()
	fmt.Println(llSlice)
}
