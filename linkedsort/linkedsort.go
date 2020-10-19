// Package linkedsort contains Daily Coding Problem #78
//
// This problem was asked by Google.
//
// Given k sorted singly linked lists, write a function to merge
// all the lists into one sorted singly linked list.
//
// Author: Andreas Atle, atle.andreas@gmail.com
package linkedsort

import (
	"container/heap"
)

// LinkedList contains a singly linked list with nodes
type LinkedList struct {
	N *Node
}

// Node contains a node with a value and a pointer to next Node
type Node struct {
	Value int
	Next  *Node
}

// NodeHeap contains a heap of Node pointers
type NodeHeap [](*Node)

// Len is an interface routine for NodeHeap
func (h NodeHeap) Len() int { return len(h) }

// Less is an interface routine for NodeHeap
func (h NodeHeap) Less(i, j int) bool { return h[i].Value < h[j].Value }

// Swap is an interface routine for NodeHeap
func (h NodeHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

// Push is an interface routine for NodeHeap
func (h *NodeHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(*Node))
}

// Pop is an interface routine for NodeHeap
func (h *NodeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// Insert inserts a value in a sorted simply linked list
func (ll *LinkedList) Insert(value int) {
	ll.insert(value, ll.N)
}

// insert is a recursive help function to Insert
func (ll *LinkedList) insert(value int, n *Node) {
	// Case where linked list is empty
	if n == nil && n == ll.N {
		ll.N = &Node{value, nil}
		return
	}

	if value <= n.Value {
		newNode := &Node{Value: n.Value, Next: n.Next}
		n.Value = value
		n.Next = newNode
		return
	}
	// Case where n is the last node in the linked list
	if n.Next == nil {
		n.Next = &Node{Value: value, Next: nil}
		return
	}

	// Recursion step
	ll.insert(value, n.Next)
}

// IsSorted returns true if the simply linked list is sorted
func (ll *LinkedList) IsSorted() bool {
	if ll == nil {
		return false
	}

	return isSorted(ll.N)
}

// isSorted is a recursive help function to IsSorted
func isSorted(n *Node) bool {
	if n == nil {
		return true
	}

	if n.Next != nil && n.Value > n.Next.Value {
		return false
	}

	return isSorted(n.Next)
}

// Merge returns a sorted simply linked list
// merged from several separe linked lists
func Merge(llSlice ...*LinkedList) *LinkedList {

	// Create a Node-heap
	nodeHeap := &NodeHeap{}
	heap.Init(nodeHeap)

	// Initialize heap with the first entry of each non-empty linked list
	for _, llCurr := range llSlice {
		if llCurr == nil || llCurr.N == nil {
			continue
		}
		heap.Push(nodeHeap, llCurr.N)
	}

	// Return if nothing to do.
	if nodeHeap.Len() == 0 {
		return &LinkedList{}
	}

	// Insert the first node from the heap into the linked list
	// and push the next in its original linked list.
	// This cumbersome initialization makes the rest of the code simpler,
	// since we can eliminate the case of an empty linked list.
	outN := heap.Pop(nodeHeap).(*Node)
	outList := &LinkedList{outN}
	if outN.Next != nil {
		heap.Push(nodeHeap, outN.Next)
	}

	for nodeHeap.Len() > 0 {
		n := heap.Pop(nodeHeap).(*Node)

		outN.Next = n
		outN = outN.Next

		if n.Next != nil {
			heap.Push(nodeHeap, n.Next)
		}
	}
	return outList
}

// Slice returns a the content of the simply linked list as a slice.
func (ll *LinkedList) Slice() []int {
	outSlice := []int{}
	for n := ll.N; n != nil; n = n.Next {
		outSlice = append(outSlice, n.Value)
	}
	return outSlice
}
